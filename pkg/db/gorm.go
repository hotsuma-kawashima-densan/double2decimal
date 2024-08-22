package db

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"gorm.io/plugin/dbresolver"
)

func NewGormDB(wCfg, rCfg Config) (*gorm.DB, func(), error) {
	cleanup := func() {}

	// ロガー設定
	gormlogger.Default = NewSlogGormLogger(
		WithLogLevel(gormlogger.Error),   // note: gorm側はすべて出力する設定。slog側でしきい値によって切り替える。
		WithSlowThreshold(8*time.Second), // note: スロークエリ判定のしきい値
	)

	// Writer接続
	wDB, err := Open(wCfg)
	if err != nil {
		return nil, cleanup, fmt.Errorf("writer接続エラー: %w", err)
	}
	wDB.SetMaxOpenConns(25)
	wDB.SetMaxIdleConns(25)
	wDB.SetConnMaxLifetime(25 * time.Second)
	cleanup = func() {
		wDB.Close()
	}

	// gorm初期化
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: wDB}))
	if err != nil {
		return nil, cleanup, fmt.Errorf("gorm初期化エラー: %w", err)
	}

	if rCfg == nil {
		slog.Info("reader設定がnilだったためwriterに接続します。")
	} else {
		// Reader接続
		rDB, err := Open(rCfg)
		if err != nil {
			return nil, cleanup, fmt.Errorf("reader接続エラー: %w", err)
		}
		rDB.SetMaxOpenConns(25)
		rDB.SetMaxIdleConns(25)
		rDB.SetConnMaxLifetime(25 * time.Second)
		cleanup = func() {
			wDB.Close()
			rDB.Close()
		}

		// リードレプリカ設定
		if err := gormDB.Use(dbresolver.
			Register(dbresolver.Config{
				Replicas:          []gorm.Dialector{mysql.New(mysql.Config{Conn: rDB})},
				TraceResolverMode: true,
			})); err != nil {
			return nil, cleanup, fmt.Errorf("gormリードレプリカ初期化エラー: %w", err)
		}
	}

	return gormDB, cleanup, nil
}

type slogGormLogger struct {
	sLogger                   *slog.Logger
	LogLevel                  gormlogger.LogLevel
	ignoreRecordNotFoundError bool
	slowThreshold             time.Duration
}

func NewSlogGormLogger(options ...LoggerOption) gormlogger.Interface {
	l := slogGormLogger{
		LogLevel:      gormlogger.Warn,
		slowThreshold: 100 * time.Millisecond,
	}
	for _, option := range options {
		option(&l)
	}
	if l.sLogger == nil {
		l.sLogger = slog.Default()
	}
	return &l
}

func (l *slogGormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func (l *slogGormLogger) Info(ctx context.Context, msg string, args ...any) {
	if l.LogLevel >= gormlogger.Info {
		l.sLogger.InfoContext(ctx, msg, args...)
	}
}

// Warn logs warn messages
func (l *slogGormLogger) Warn(ctx context.Context, msg string, args ...any) {
	if l.LogLevel >= gormlogger.Warn {
		l.sLogger.WarnContext(ctx, msg, args...)
	}
}

// Error logs error messages
func (l *slogGormLogger) Error(ctx context.Context, msg string, args ...any) {
	if l.LogLevel >= gormlogger.Error {
		l.sLogger.ErrorContext(ctx, msg, args...)
	}
}

func (l *slogGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= gormlogger.Silent {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error && (!l.ignoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		l.sLogger.Log(ctx, slog.LevelError, err.Error(),
			slog.Any("cause", err),
			slog.String("sql", sql),
			slog.Duration("elapsed", elapsed),
			slog.Int64("rows", rows),
			slog.String("file", utils.FileWithLineNum()),
		)
	case l.LogLevel >= gormlogger.Warn && l.slowThreshold != 0 && elapsed > l.slowThreshold:
		sql, rows := fc()
		l.sLogger.Log(ctx, slog.LevelWarn, fmt.Sprintf("slow sql query [%s >= %v]", elapsed, l.slowThreshold),
			slog.Bool("slow_query", true),
			slog.String("sql", sql),
			slog.Duration("elapsed", elapsed),
			slog.Int64("rows", rows),
			slog.String("file", utils.FileWithLineNum()),
		)
	case l.LogLevel >= gormlogger.Info:
		sql, rows := fc()
		l.sLogger.Log(ctx, slog.LevelInfo, fmt.Sprintf("SQL query executed [%s]", elapsed),
			slog.String("sql", sql),
			slog.Duration("elapsed", elapsed),
			slog.Int64("rows", rows),
			slog.String("file", utils.FileWithLineNum()),
		)
	}
}

type LoggerOption func(l *slogGormLogger)

func WithLogger(log *slog.Logger) LoggerOption {
	return func(l *slogGormLogger) {
		l.sLogger = log
	}
}

func WithLogLevel(level gormlogger.LogLevel) LoggerOption {
	return func(l *slogGormLogger) {
		l.LogLevel = level
	}
}

func WithSlowThreshold(threshold time.Duration) LoggerOption {
	return func(l *slogGormLogger) {
		l.slowThreshold = threshold
	}
}

func WithRecordNotFoundError() LoggerOption {
	return func(l *slogGormLogger) {
		l.ignoreRecordNotFoundError = false
	}
}
