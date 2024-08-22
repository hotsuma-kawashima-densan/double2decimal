package db

import (
	"crypto/tls"
	"crypto/x509"
	_ "embed"
	"errors"
	"fmt"
	"net/url"

	"github.com/go-sql-driver/mysql"
)

var (
	//go:embed global-bundle.pem
	rdsPEM []byte
)

// MySqlConfig MySQL設定
type MySqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DbName   string `json:"db_name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (_ *MySqlConfig) DriverName() string {
	return "mysql"
}

func (c *MySqlConfig) DSN() (string, error) {
	opt, _ := url.ParseQuery("")
	opt.Set("parseTime", "true")
	opt.Set("loc", "Asia/Tokyo")

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		c.User, c.Password, c.Host, c.Port, c.DbName, opt.Encode(),
	), nil
}

// AuroraMySQLConfig Aurora MySQL設定
type AuroraMySQLConfig struct {
	Endpoint string `json:"endpoint"`
	Port     int    `json:"port"`
	DbName   string `json:"db_name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (*AuroraMySQLConfig) DriverName() string {
	return "mysql"
}

func (c *AuroraMySQLConfig) DSN() (string, error) {
	pool := x509.NewCertPool()
	if ok := pool.AppendCertsFromPEM(rdsPEM); !ok {
		return "", errors.New("RDS TLS証明書読込エラー")
	}

	const tlsKey = "rds"
	if err := mysql.RegisterTLSConfig(tlsKey, &tls.Config{
		RootCAs: pool,
	}); err != nil {
		return "", fmt.Errorf("RDS TLS証明書登録エラー: %w", err)
	}

	opt, _ := url.ParseQuery("")
	opt.Set("tls", tlsKey)
	opt.Set("parseTime", "true")
	opt.Set("loc", "Asia/Tokyo")

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		c.User, c.Password, c.Endpoint, c.Port, c.DbName, opt.Encode(),
	), nil
}
