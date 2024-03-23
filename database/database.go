package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "os"
)

var db *gorm.DB

func StartDB() {
    // Read DATABASE_URL from environment variable
    dbURL := os.Getenv("DATABASE_URL")

    // Open database connection
    var err error
    db, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{})
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }

    log.Println("Connected to the database")
}

func GetDB() *gorm.DB {
    return db
}
