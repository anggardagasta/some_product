package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	DBEngine = "DB_ENGINE"
	DBHost   = "DB_HOST"
	DBPort   = "DB_PORT"
	DBUser   = "DB_USER"
	DBPwd    = "DB_PWD"
	DBName   = "DB_NAME"
)

func Env(key string, def string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return def
}

func (c *Config) InitMysql() error {
	driver := Env(DBEngine, "mysql")
	user := Env(DBUser, "root")
	password := Env(DBPwd, "anggarda")
	host := Env(DBHost, "127.0.0.1")
	port := Env(DBPort, "3306")
	dbName := Env(DBName, "mini_wallet")

	dataSource := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true"

	db, err := sql.Open(driver, dataSource)

	if err != nil {
		return err
	}

	c.DB = db

	return nil
}
