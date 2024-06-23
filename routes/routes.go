package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/accounts", createAccount)
	server.GET("/accounts/:id", getAccount)
	server.GET("/accounts", listAccounts)
	server.PUT("/account", updateAccount)
	server.DELETE("/accounts/:id", deleteAccount)

	server.POST("/entry", createEntry)
	server.GET("/entries/:account_id", listEntries)

	server.POST("/transfer", createTransfer)
	server.GET("/transfers/:account_id", listTransfers)
}
