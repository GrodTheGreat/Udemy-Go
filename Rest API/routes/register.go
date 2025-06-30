package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rest/models"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id, could not parse id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event. Try again later"})
		log.Println(err)
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register. Try again later"})
		log.Println(err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully registered for event."})
}

func cancelRegistration(context *gin.Context) {}
