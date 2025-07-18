package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rest/models"
	"strconv"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch events. Try again later"})
		log.Println(err)
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		log.Println(err)
		return
	}

	userId := context.GetInt64("userId")

	event.UserID = userId
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
		log.Println(err)
		return
	}

	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id, could not parse id"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event. Try again later"})
		log.Println(err)
		return
	}

	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to update this event."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		log.Println(err)
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update event. Try again later"})
		log.Println(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated!", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id, could not parse id"})
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event. Try again later"})
		log.Println(err)
		return
	}

	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to delete this event."})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete event. Try again later"})
		log.Println(err)
		return
	}

	context.JSON(http.StatusNoContent, nil)
}
