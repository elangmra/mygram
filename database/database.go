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
	// Konfigurasi koneksi MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	// Buka koneksi ke MySQL
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	// Auto migrate tabel-tabel model
	db.Debug().AutoMigrate(&models.User{}, &models.Comment{}, &models.Photo{}, &models.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
