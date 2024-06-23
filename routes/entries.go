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

func createEntry(ctx *gin.Context) {
	body := dto.CreateEntryParamets{}
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

	err = db.CreateEntry(body)

	if err != nil {
		log.Println("ERROR createEntry: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't create the new Entry.")
		return
	}

	ctx.JSON(http.StatusOK, "The entry was successfully created.")
}

func listEntries(ctx *gin.Context) {
	accountId, err := strconv.ParseInt(ctx.Param("account_id"), 10, 64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	entries, err := db.ListEntries(accountId)

	if err != nil {
		log.Println("ERROR listEntry: accountId: ", accountId, err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, "Couldn't get the entry")
		return
	}

	ctx.JSON(http.StatusOK, entries)
}
