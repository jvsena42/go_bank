package db

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const createAccountQuery = `
INSERT INTO accounts (
	owner,
	balance,
	currency
) VALUES (
	$1, $2, $3
)
`

type CreateAccountParameters struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func CreateAccount(ctx *gin.Context) {
	body := CreateAccountParameters{}
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

	_, err = Db.Exec(createAccountQuery, body.Owner, body.Balance, body.Currency)

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
