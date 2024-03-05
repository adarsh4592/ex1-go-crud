package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	// dsn := fmt.Sprintf("host=%s port=%s dbuser=%s dbname=%s password=%s sslmode=disable",
	// 	os.Getenv("dbUser"), os.Getenv("password"), os.Getenv("dbName"), os.Getenv("port"), os.Getenv("host"))
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = db

	if err != nil {
		log.Fatal(" Db connection failed ")
	}

}
