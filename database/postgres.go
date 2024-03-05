package database

import (
	"log"
	"os"
  
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB
var err error

type Student struct {
	gorm.Model
	Name  string `json:"name"`
	Class string `json:"class"`
}

func LoadEnvVariables(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

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