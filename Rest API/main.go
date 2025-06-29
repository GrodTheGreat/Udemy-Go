package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rest/db"
	"rest/models"
	"strconv"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.POST("/events", createEvent)
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	log.Fatal(server.Run(":8080"))
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch events. Try again later"})
		log.Fatal(err)
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		log.Fatal(err)
		return
	}

	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save event. Try again later"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id, could not parse id"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event. Try again later"})
		log.Fatal(err)
		return
	}

	context.JSON(http.StatusOK, event)
}
