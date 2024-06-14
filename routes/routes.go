package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jvsena42/go_bank/db"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/register", db.CreateAccount)
}
