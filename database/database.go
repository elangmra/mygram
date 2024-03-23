package database

import (
	"fmt"
	"log"
	"mygram/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 3306 // Port MySQL default adalah 3306
	user     = "root"
	password = "rizki121"          // Ganti dengan kata sandi MySQL Anda
	dbname   = "mygram" // Ganti dengan nama database Anda
	db       *gorm.DB
	err      error
)

func StartDB() {
    // Buka koneksi ke MySQL
    db, err = gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect database: ", err)
    }

    // Auto migrate tabel-tabel model
    db.Debug().AutoMigrate(&models.User{}, &models.Comment{}, &models.Photo{}, &models.SocialMedia{})
}


func GetDB() *gorm.DB {
	return db
}
