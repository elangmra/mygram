package database

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func StartDB() {
    // Read DATABASE_URL from environment variable
    dbURL := os.Getenv("DATABASE_URL")

    // Open database connection
    var err error
    db, err = sql.Open("mysql", dbURL)
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }

    // Ping the database to verify the connection
    err = db.Ping()
    if err != nil {
        log.Fatal("Error pinging the database:", err)
    }

    log.Println("Connected to the database")
}

func GetDB() *sql.DB {
    return db
}
