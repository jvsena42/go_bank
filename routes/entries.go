package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jvsena42/go_bank/db"
	"github.com/jvsena42/go_bank/dto"
)

func createEntry(ctx *gin.Context) {
	body := dto.CreateAccountParameters{}
	data, err := ctx.GetRawData()
	if err != nil {
		log.Println("ERROR createEntry: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Error finding entry")
		return
	}

	err = json.Unmarshal(data, &body)
	if err != nil {
		log.Println("ERROR createEntry: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Bad Input")
		return
	}

	err = db.CreateAccount(body)

	if err != nil {
		log.Println("ERROR createEntry: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't create the new Entry.")
		return
	}

	ctx.JSON(http.StatusOK, "The entry was successfully created.")
}
