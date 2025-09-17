package routes

import (
	"RESTApi/models"
	//"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	allEvents, err := models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "unable to get all events"})
	}
	context.JSON(http.StatusOK, gin.H{"Events retrieved successfully": allEvents})

}

func CreateEvents(context *gin.Context) {
	var event models.Events

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad request"})
		return
	}

	event.ID = 1

	err = event.Save()
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad request"})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{"Created event": event})
}