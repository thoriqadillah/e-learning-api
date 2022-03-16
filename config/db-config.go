package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

//For creating database connection
func SetupDB() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic(errEnv.Error())
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Error connecting database : " + err.Error())
	}

	DB = db
}

//For closing database connection
func CloseDB() {
	dbCon, err := DB.DB()
	if err != nil {
		panic("Error closing database : " + err.Error())
	}

	dbCon.Close()
}
