package helper

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
	"blog-go-gin/config"
)

var DB *sqlx.DB
func InitDb() error {
	var err error
	cfg := config.GetCfg()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DbUsername, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbDatabase)
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}

func GetDb() *sqlx.DB {
	return DB
}