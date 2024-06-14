package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jvsena42/go_bank/db"
	"github.com/jvsena42/go_bank/routes"
)

func main() {
	db.ConnectDatabase()
	route := gin.Default()
	routes.RegisterRoutes(route)
	route.Run(":8080")
}
