package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var mysqlAddress = os.Getenv("MYSQL_ADDRESS")
var mysqlRootPassword = os.Getenv("MYSQL_ROOT_PASSWORD")

// Connect to mysql with gorm
func Connect() (*gorm.DB, error) {
	dsn := "root:" + mysqlRootPassword + "@tcp(" + mysqlAddress + ")/LinkShortener?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}

func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		return err
	}
	return nil
}
