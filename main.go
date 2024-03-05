package main

import (
	"github.com/adarsh4592/go-crud/controllers"
	"github.com/adarsh4592/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	// initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "student db operation exercise",
		})
	})

	r.GET("/stud/:id", controllers.ReadStud)
	r.GET("/stud", controllers.ReadStuds)
	r.POST("/stud", controllers.CreateStud)
	r.PUT("/stud/:id", controllers.UpdateStud)
	r.DELETE("/stud/:id", controllers.DeleteStud)
	r.Run()
}
