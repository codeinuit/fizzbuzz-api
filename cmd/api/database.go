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

	err = db.AutoMigrate(&models.Stats{})
	if err != nil {
		return &Database{}, err
	}

	return &Database{
		db: db,
	}, err
}

func (db *Database) Count(i any) int64 {
	var res int64

	_ = db.db.Where(i).Count(&res)
	return res
}

func (db *Database) CountUsage() (models.Stats, error) {
	var result []models.Stats

	db.db.Model(&models.Stats{}).Find(&result)

	var mostUsed models.Stats
	for _, stat := range result {
		if stat.Use > mostUsed.Use {
			mostUsed = stat
		}
	}

	return mostUsed, nil
}

func (db *Database) UsageUpdate(m models.Stats) {
	db.db.Transaction(func(tx *gorm.DB) error {
		var result models.Stats

		err := tx.Where(m).Take(&result).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			fmt.Println(err.Error())
			return err
		} else if err == gorm.ErrRecordNotFound {
			fmt.Println("creating entry")
			m.Use = 1
			err = tx.Create(&m).Error
			return err
		}

		fmt.Println("updating", result.Use)
		err = tx.Model(&result).Update("use", result.Use+1).Error
		return err

	})
}
