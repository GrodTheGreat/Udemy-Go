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
		log.Println(err)
		return
	}

	var user models.User
	user.Email = request.Email

	err = user.Save(request.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		log.Println(err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func login(context *gin.Context) {
	var request LoginRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse body"})
		log.Println(err)
		return
	}

	var user models.User
	user.Email = request.Email

	err = user.ValidateCredentials(request.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		log.Println(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User login successful"})

}
