package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/accounts", createAccount)
	server.GET("/accounts/:id", getAccount)
	server.GET("/accounts", listAccounts)
	//server.PUT("/account", updateAccount)
	//server.GET("/accounts/:id", deleteAccount)
}
