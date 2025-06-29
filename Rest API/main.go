package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"rest/db"
	"rest/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	log.Fatal(server.Run(":8080"))
}
