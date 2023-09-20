package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbHost := os.Getenv("dbHost")
	dbUser := os.Getenv("dbUser")
	dbPass := os.Getenv("dbPass")
	dbName := os.Getenv("dbName")
	dbPort := os.Getenv("dbPort")
	dbSSLMode := os.Getenv("dbSSLMode")
	dbTimezone := os.Getenv("dbTimezone")

	if len(dbSSLMode) == 0 {
		dbSSLMode = "disable"
	}

	dsn := "dbHost=" + dbHost +
		" dbUser=" + dbUser +
		" dbPassword=" + dbPass +
		" db_dbname=" + dbName +
		" dbPort=" + dbPort +
		" dbSSLMode=" + dbSSLMode +
		" dbTimezone=" + dbTimezone

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Album{})

	DB = database
}
