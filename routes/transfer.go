package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jvsena42/go_bank/db"
	"github.com/jvsena42/go_bank/dto"
)

func createTranfer(ctx *gin.Context) {
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
