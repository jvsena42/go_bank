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
		ctx.JSON(http.StatusOK, "User is successfully created.")
	}

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't create the new Account.")
		log.Println("ERROR createAccount: ", err)
		return
	} else {
		ctx.JSON(http.StatusCreated, "User is Account created.")
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
	} else {
		ctx.JSON(http.StatusOK, "User is successfully created.")
	}

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't get the account")
	}

	ctx.JSON(http.StatusOK, account)
}

func listAccounts(ctx *gin.Context) {

	accounts, err := db.ListAccounts()

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, "Couldn't create the new Account.")
	} else {
		ctx.JSON(http.StatusOK, "User is successfully created.")
	}

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't create the new Account.")
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't update the new Account.")
	} else {
		ctx.JSON(http.StatusOK, "Account is updated created.")
	}

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't update the Account.")
	} else {
		ctx.JSON(http.StatusCreated, "Account is updated created.")
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
	} else {
		ctx.JSON(http.StatusNoContent, "Account was successfully deleted.")
	}

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't delete the account")
	}

	ctx.JSON(http.StatusNoContent, "Account was successfully deleted.")
}
