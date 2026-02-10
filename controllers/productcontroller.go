package controllers

import (
	
	"myproject/models"
	"myproject/db"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context){
	var prod models.Product
	if err := c.ShouldBindJSON(&prod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&prod)
	c.JSON(http.StatusOK, gin.H{"message": "Created a product successfully"})
}

func GetProducts(c *gin.Context) {
	var prod []models.Product
	if err := db.DB.Find(&prod).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusNotFound, prod)
}