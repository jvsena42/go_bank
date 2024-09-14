package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jvsena42/go_bank/db/sqlc"
)

type CreateAccountRequest struct {
	Owner    string `json:"owner" bindin:"required"`
	Currency string `json:"currency" bindin:"required, oneof= USD EUR BTC"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, account)
}
