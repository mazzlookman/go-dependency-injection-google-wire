package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-salaries-app/helper"
	"time"
)

func NewDBTest() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/salaries_app_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/salaries_app")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
