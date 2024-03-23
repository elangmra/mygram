// main.go

package main

import (
    "mygram/database"
    "mygram/router"
    "os"
)

func main() {
    // Start database connection
    database.StartDB()

    // Start router
    r := router.StartDB()
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}
