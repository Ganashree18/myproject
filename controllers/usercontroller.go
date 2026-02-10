package controllers

import (
	"fmt"
	"myproject/db"
	"myproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Sample1() {
	fmt.Println("Sampleeeee")
}

func Createee(c *gin.Context) {
	var prod models.Product
	if err := c.ShouldBindJSON(&prod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&prod)
	c.JSON(http.StatusOK, gin.H{"message": "Created a product successfully"})
}
