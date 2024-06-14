package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jvsena42/go_bank/db"
)

func main() {
	route := gin.Default()

	route.Run(":8080")
	db.ConnectDatabase()

}
