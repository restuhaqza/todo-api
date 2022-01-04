package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

func Connect(config DBConfig) *sql.DB {
	db, err := sql.Open(config.Driver, fmt.Sprintf("%s:%s@%s/%s", config.User, config.Password, config.Host, config.DB))

	if err != nil {
		log.Fatal(err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("Database Connected")

	return db
}
