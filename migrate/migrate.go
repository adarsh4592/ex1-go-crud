package main

import (
	"github.com/adarsh4592/go-crud/initializers"
	"github.com/adarsh4592/go-crud/models"
)

func init() {

	initializers.ConnectToDB()
	initializers.LoadEnvVariables()
}

func main() {
	initializers.DB.AutoMigrate(&models.Student{})
}
