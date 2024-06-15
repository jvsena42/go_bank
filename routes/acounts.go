package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

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
		ctx.AbortWithStatusJSON(400, "Bad Input")
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
