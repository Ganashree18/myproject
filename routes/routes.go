package routes

import (
	"github.com/gin-gonic/gin"
	"myproject/controllers"

)

func SetRouter() *gin.Engine{
	r:= gin.Default()
	r.POST("/create-prod/",controllers.CreateProduct)
	r.GET("/get-prod/",controllers.GetProducts)

	
	return r

}
