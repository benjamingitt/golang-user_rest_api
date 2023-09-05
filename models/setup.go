package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=localhost user=postgres dbname=postgres port=5432 password=postgres sslmode=disable"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})  

    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&User{},&Picture{})  

    DB = database
}