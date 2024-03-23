package main

import (
	"mygram/database"
	"mygram/router"
)

func main() {
	database.StartDB()

	r := router.StartDB()
	r.Run(":8080")
}
