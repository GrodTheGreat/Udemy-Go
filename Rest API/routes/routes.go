package routes

import (
	"github.com/gin-gonic/gin"
	"rest/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)

	authenticated := server.Group("/events")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("", createEvent)
	authenticated.PUT("/:id", updateEvent)
	authenticated.DELETE("/:id", deleteEvent)

}
