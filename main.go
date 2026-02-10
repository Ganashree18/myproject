package main

import (
	"log"
	
	"myproject/db"
	"myproject/models"
	"myproject/routes"
)


func main(){
	db.ConnectPostgres()

	err:= db.DB.AutoMigrate(&models.Product{})
	if err != nil{
		log.Fatal("migration unsuccessful", err)
	}
	r:= routes.SetRouter()
	r.Run(":8080")

}










		
	
	

