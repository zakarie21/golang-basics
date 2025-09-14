package routes

import "github.com/gin-gonic/gin"

func RouterInitialisation(server *gin.Engine) {
	server.Handle("GET", "/events", getEvents)
	server.Handle("POST", "/events", CreateEvents)
	
}
