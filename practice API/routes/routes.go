package routes

import "github.com/gin-gonic/gin"

func RouterInitialisation(server *gin.Engine) {
	server.Handle("GET", "/events", getEvents)
	server.Handle("POST", "/events", CreateEvents)
	server.Handle("GET", "/events/:id", getAnEvent)
	server.Handle("DELETE", "/events/:id", deleteAnEvent)
	server.Handle("PUT", "/events/:id", updateAnEvent)
	server.Handle("POST", "/signup", signUp)
	

}
