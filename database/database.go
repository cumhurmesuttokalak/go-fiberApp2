package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("gofiberApp2"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
		os.Exit(2)
	}
	log.Fatal("connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations...")
	//TODO: add migrations

	Database = DbInstance{Db: db}
}
