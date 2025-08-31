package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"fmt"
	"RESTApi/models"
) 




func main() {
	server := gin.Default()

	server.Handle("GET", "/events", getEvents)
	server.Handle("POST", "/events", CreateEvents)
	

	server.Run("localhost:8080")
}

func getEvents(context *gin.Context) {
	allEvents:= models.GetAllEvents()
	context.JSON(http.StatusOK, gin.H{"Events retrieved successfully": allEvents})

}


func CreateEvents(context *gin.Context) {
	var event models.Events
	
	err:= context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad request"})
		return
	}

	event.ID = 1

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"Created event" : event})
}

