package routes

import (
	"RESTApi/models"
	
	"strconv"

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

func getAnEvent(context *gin.Context) {
	eventID,err:= strconv.ParseInt(context.Param("id"),10,64)

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID passed"})
		return 
	}
	getOneEvent, err:= models.GetEvent(eventID)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "No Search event present"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Event retrieved successfully": getOneEvent})
}

func CreateEvents(context *gin.Context) {
	var event models.Events

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad request"})
		return
	}

	

	err = event.Save()
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad request"})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{"Created event": event})
}

