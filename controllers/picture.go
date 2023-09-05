package controllers

import (
	"log"
	"net/http"
	"rest/models"

	"github.com/gin-gonic/gin"
)

type CreatePictureInput struct {
    Src string `json:"src" binding:"required"`
    UserID uint `json:"user_id" binding:"required"`
}

func CreatePicture(c *gin.Context) {
    var input CreatePictureInput
    if err := c.ShouldBindJSON(&input); err != nil {
		log.Fatalf("impossible to start server: %s", err)
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    picture := models.Picture{Src: input.Src, UserID: input.UserID}
    if result := models.DB.Create(&picture); result.Error != nil {
		log.Fatalf("impossible to start server: %s", result.Error)
        c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": result.Error})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": picture})
}