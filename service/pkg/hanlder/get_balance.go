package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"service/pkg/model"
	"service/pkg/response"
	"strconv"
)

func (h *Handler) getBalance(ctx *gin.Context) {
	var (
		account model.Account
		resp    response.Response
		err     error
	)

	resp.RespWriter = ctx.Writer

	param := ctx.Param("user_id")
	id, _ := strconv.Atoi(param)
	account.UserID = int64(id)

	account, err = h.service.Payment.GetAccount(account.UserID)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			resp.SetStatusBadRequest()
			log.Printf("Database don't have informantion about account with id=%d. err: %s", account.UserID, err)
			return
		} else {
			resp.SetStatusInternalServerError()
			log.Printf("Informantion about account hasn't been gotten. err: %s", err)
			return
		}
	}

	accountToJSON, err := json.Marshal(account)
	if err != nil {
		resp.SetStatusInternalServerError()
		log.Println("Data hasn't been encoded to JSON!")
		return
	}

	resp.SetStatusOk()
	resp.SetData(accountToJSON)
}
