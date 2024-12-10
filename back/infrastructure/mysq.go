package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

/*
* DB接続用
 */
func NewDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// MySqlの接続設定
	config := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    os.Getenv("DB_NET"),
		Addr:   fmt.Sprintf("%v:%v", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName: os.Getenv("DB_NAME"),
	}

	// DBとの接続開始
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	// DB疎通確認
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
