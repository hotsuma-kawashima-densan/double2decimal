package main

import (
	"project+test/pkg/db"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// ソースを生成する元になるDBへの接続
	localDB, err := db.Open(&db.MySqlConfig{
		Host:     "127.0.0.1",
		Port:     13306,
		DbName:   "test",
		User:     "root",
		Password: "root",
	})

	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: localDB}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 全体設定
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./pkg/db/query",
		FieldNullable: true,
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(db)

	// 生成対象のテーブル指定とテーブル個別の設定
	g.ApplyBasic(
		g.GenerateAllTable()...,
	)

	g.Execute()
}
