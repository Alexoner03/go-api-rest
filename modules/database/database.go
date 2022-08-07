package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func Connect() error {

	ConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	var err error

	Connection, err = gorm.Open(mysql.Open(ConnectionString), &gorm.Config{})

	return err
}

func SetModelToMigrate(models ...interface{}) error {
	err := Connection.AutoMigrate(models...)
	return err
}
