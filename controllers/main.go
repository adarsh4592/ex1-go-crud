package controllers
 
import (
   "errors"
   "net/http"

   "github.com/adarsh4592/go-crud/database"
   "github.com/gin-gonic/gin"
)

func CreateStud(c *gin.Context) {
	var stud *database.Student
	err := c.ShouldBind(&stud)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(stud)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a student",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"stud": stud,
	})
	return
 }
 
 func ReadStud(c *gin.Context) {
	var stud *database.Student
	id := c.Param("id")
	res := database.DB.Find(&stud, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "student not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"stud": stud,
	})
	return
 }

 func ReadStuds(c *gin.Context) {
	var stud []database.Student
	res := database.DB.Find(&stud)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("student not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"stud": stud,
	})
	return
 }

 func UpdateStud(c *gin.Context) {
	var stud *database.Student
	id := c.Param("id")
	err := c.ShouldBind(&stud)
   
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
  
	var updateStud *database.Student
	res := database.DB.Model(&updateStud).Where("id = ?", id).Updates(stud)
  
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "student not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"stud": stud,
	})
	return
 }
 

 func DeleteStud(c *gin.Context) {
	var stud *database.Student
	id := c.Param("id")
	res := database.DB.Find(&stud, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "student not found",
		})
		return
	}
	database.DB.Delete(&stud)
	c.JSON(http.StatusOK, gin.H{
		"message": "student deleted successfully",
	})
	return
 }
 