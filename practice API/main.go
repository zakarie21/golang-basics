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

	server.Run("localhost:8080")
}

func getEvents(context *gin.Context) {
	allEvents:= models.GetAllEvents()
	context.JSON(http.StatusOK, gin.H{"Events retrieved successfully": allEvents})

}