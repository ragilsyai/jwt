package main

import (
	"jwt/database"
	"jwt/router"
)

func main() {
	database.StartDB()
	r := router.StarApp()
	r.Run(":8080")
}
