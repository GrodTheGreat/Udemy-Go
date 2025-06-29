package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rest/models"
)

func signup(context *gin.Context) {
	var request models.SignupRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse body"})
		log.Fatal(err)
		return
	}

	var user models.User
	user.Email = request.Email

	err = user.Save(request.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		log.Fatal(err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
