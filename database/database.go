package database

import (
	"fmt"
	"log"
	"os"

	"github.com/cumhurmesuttokalak/go-fiberApp2/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("dbgofiberApp2"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
		os.Exit(2)
	}
	fmt.Println("connected to the database successfully")

	fmt.Println("Running Migrations...")
	//TODO: add migrations
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	Database = DbInstance{Db: db}
}
