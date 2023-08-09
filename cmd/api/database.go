package main

import (
	"fmt"
	"os"

	"github.com/codeinuit/fizzbuzz-api/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MYSQL_HOST = "MYSQL_HOST"
	MYSQL_DB   = "MYSQL_DB"
	MYSQL_USER = "MYSQL_USER"
	MYSQL_PASS = "MYSQL_PASS"
	MYSQL_PORT = "MYSQL_PORT"
)

type Database struct {
	db *gorm.DB
}

func InitDatabase() (*Database, error) {
	dbHost := os.Getenv(MYSQL_HOST)
	dbDatabase := os.Getenv(MYSQL_DB)
	dbUser := os.Getenv(MYSQL_USER)
	dbPass := os.Getenv(MYSQL_PASS)
	dbPort := os.Getenv(MYSQL_PORT)

	str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbDatabase)
	fmt.Println(str)

	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		return &Database{}, err
	}

	err = db.AutoMigrate(models.FizzBuzzRequestStat{})
	if err != nil {
		return &Database{}, err
	}

	return &Database{
		db: db,
	}, err
}
