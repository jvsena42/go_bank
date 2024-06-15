package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jvsena42/go_bank/db"
	"github.com/jvsena42/go_bank/dto"
)

func createAccount(ctx *gin.Context) {
	body := dto.CreateAccountParameters{}
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "User is not defined")
		return
	}

	err = json.Unmarshal(data, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Bad Input")
		return
	}

	err = db.CreateAccount(body)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't create the new Account.")
	} else {
		ctx.JSON(http.StatusOK, "User is successfully created.")
	}

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't create the new Account.")
	} else {
		ctx.JSON(http.StatusOK, "User is Account created.")
	}
}

func getAccount(ctx *gin.Context) {
	accountId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = db.GetAccount(accountId)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, "Couldn't create the new Account.")
	} else {
		ctx.JSON(http.StatusOK, "User is successfully created.")
	}

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't create the new Account.")
	} else {
		ctx.JSON(http.StatusOK, "User is Account created.")
	}
}
