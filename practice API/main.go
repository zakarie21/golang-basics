package main

import (
	"github.com/gin-gonic/gin"
	"RESTApi/routes"
) 




func main() {
	server := gin.Default()

	routes.RouterInitialisation(server)
	

	server.Run("localhost:8080")
}



