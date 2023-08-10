package mysql

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

type MySQLDatabase struct {
	db *gorm.DB
}

// InitDatabase is used to return the MySQL database instance
func InitDatabase() (*MySQLDatabase, error) {
	dbHost := os.Getenv(MYSQL_HOST)
	dbDatabase := os.Getenv(MYSQL_DB)
	dbUser := os.Getenv(MYSQL_USER)
	dbPass := os.Getenv(MYSQL_PASS)
	dbPort := os.Getenv(MYSQL_PORT)

	str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbDatabase)

	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		return &MySQLDatabase{}, err
	}

	err = db.AutoMigrate(&models.Stats{})
	if err != nil {
		return &MySQLDatabase{}, err
	}

	return &MySQLDatabase{
		db: db,
	}, err
}

// CountUsage is used to return the number of usages of FizzBuzz from the DB
func (db *MySQLDatabase) CountUsage() (models.Stats, error) {
	var result []models.Stats

	err := db.db.Model(&models.Stats{}).Find(&result).Error

	// @TODO: use max() instead
	// reason of the comment : encountered a library bug with the driver
	// getting every entries is usually a bad practice
	var mostUsed models.Stats
	for _, stat := range result {
		if stat.Use > mostUsed.Use {
			mostUsed = stat
		}
	}

	return mostUsed, err
}

// UsageUpdate is used to increment the number of usages for FizzBuzz endpoint
// to the database
func (db *MySQLDatabase) UsageUpdate(m models.Stats) error {
	return db.db.Transaction(func(tx *gorm.DB) error {
		var result models.Stats

		err := tx.Where(m).Take(&result).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			fmt.Println(err.Error())
			return err
		} else if err == gorm.ErrRecordNotFound {
			m.Use = 1
			err = tx.Create(&m).Error
			return err
		}

		err = tx.Model(&result).Update("use", result.Use+1).Error
		return err
	})
}
