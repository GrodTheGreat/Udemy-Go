package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)

	log.Fatal(server.Run(":8080"))
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
