package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "myapi/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&models.User{})

    DB = database
}
