package main

import (
	"log"
	"net/http"
	"rest/controllers"
	"rest/models"

	"github.com/gin-gonic/gin"
)

func main() {
	
	r := gin.Default()
	models.ConnectDatabase() 

	r.POST("/users", controllers.CreateUser)
	r.POST("/picture", controllers.CreatePicture)
	r.GET("/users", controllers.FindUsers)
	r.GET("/users/:id", controllers.FindUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	err := r.Run(":8083")
	if err != nil {
	   log.Fatalf("impossible to start server: %s", err)
	}

}