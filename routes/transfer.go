package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jvsena42/go_bank/db"
	"github.com/jvsena42/go_bank/dto"
)

func createTransfer(ctx *gin.Context) {
	body := dto.CreateTransferParameters{}
	data, err := ctx.GetRawData()
	if err != nil {
		log.Println("ERROR createTransfer: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Error finding transfer")
		return
	}

	err = json.Unmarshal(data, &body)
	if err != nil {
		log.Println("ERROR createTransfer: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Bad Input")
		return
	}

	err = db.CreateTransfer(body)

	if err != nil {
		log.Println("ERROR createTransfer: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't create the new transfer.")
		return
	}

	ctx.JSON(http.StatusOK, "The transfer was successfully created.")
}

func listTransfers(ctx *gin.Context) {
	accountId, err := strconv.ParseInt(ctx.Param("account_id"), 10, 64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	transfers, err := db.ListTransfers(accountId)

	if err != nil {
		log.Println("ERROR listTransfer: accountId: ", accountId, err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, "Couldn't find the transfer")
		return
	}

	ctx.JSON(http.StatusOK, transfers)
}
