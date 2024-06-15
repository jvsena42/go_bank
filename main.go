package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jvsena42/go_bank/db"
	"github.com/jvsena42/go_bank/routes"
)

func main() {
	db.ConnectDatabase()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
