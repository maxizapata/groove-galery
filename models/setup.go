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
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

	dsn := "host=" + dbHost +
		" user=" + dbUser +
		" password=" + dbPass +
		" dbname=" + dbName +
		" port=" + dbPort +
		" sslmode=" + dbSSLMode +
		" TimeZone=" + dbTimezone

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Album{})

	DB = database
}
