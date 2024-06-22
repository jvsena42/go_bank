package routes

import (
	"encoding/json"
	"fmt"
	"log"
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
		log.Println("ERROR createAccount: ", err)
		return
	} else {
		ctx.JSON(http.StatusOK, "Account was successfully created.")
	}
}

func getAccount(ctx *gin.Context) {
	accountId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	account, err := db.GetAccount(accountId)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, "Couldn't get the account")
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func listAccounts(ctx *gin.Context) {

	accounts, err := db.ListAccounts()

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, "Couldn't get the Accounts.")
	}

	ctx.JSON(http.StatusOK, accounts)
}

func updateAccount(ctx *gin.Context) {
	body := dto.UpdateAccountParameters{}
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

	err = db.UpdateAccount(body)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't update the Account.")
	} else {
		ctx.JSON(http.StatusOK, "The account was Successfuly updated")
	}
}

func deleteAccount(ctx *gin.Context) {
	accountId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = db.DeleteAccount(accountId)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, "Couldn't find the account")
		return
	}

	ctx.JSON(http.StatusNoContent, "Account was successfully deleted.")
}
