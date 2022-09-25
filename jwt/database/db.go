package database

import (
	"fmt"
	"jwt/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "location"
	user     = "postgres"
	password = "postgres"
	dbPort   = "4444"
	dbname   = "Jwt"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprint("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to databasee")
	}

	fmt.Println("sukses koneksi ke database")
	db.Debug().AutoMigrate(models.User{}, models.Product{})

}

func GetDB() *gorm.DB {
	return db
}
