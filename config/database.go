package config

import (
	"fmt"
	"log"

	models "github.com/apriliantocecep/ordent-restapi-artikel/app/Models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseHandler struct {
	DB *gorm.DB
}

func NewDB(config *Config) *gorm.DB {
	handler := InitDatabaseHandler(config)
	return handler.DB
}

func InitDatabaseHandler(config *Config) DatabaseHandler {
	var db *gorm.DB

	if config.DB_CONNECTION == "pgsql" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
			config.DB_HOST,
			config.DB_USERNAME,
			config.DB_PASSWORD,
			config.DB_DATABASE,
			config.DB_PORT,
		)
		database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		db = database
	}

	if config.DB_CONNECTION == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.DB_USERNAME,
			config.DB_PASSWORD,
			config.DB_HOST,
			config.DB_PORT,
			config.DB_DATABASE,
		)
		database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		db = database
	}

	// auto migrate
	db.AutoMigrate(&models.User{}, &models.Article{})

	return DatabaseHandler{
		DB: db,
	}
}
