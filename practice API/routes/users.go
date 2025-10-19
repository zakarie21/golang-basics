package routes

import (
	"RESTApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.Users

	err:= context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad Request"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Internal server error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Message" : "User signed up successfully"})
}