package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"service/pkg/model"
	"service/pkg/response"
)

func (h *Handler) createAccount(ctx *gin.Context) {
	var (
		accountID model.ID
		resp      response.Response
		err       error
	)
	resp.RespWriter = ctx.Writer

	accountID.UserID, err = h.service.Payment.CreateAccount()
	if err != nil {
		resp.SetStatusInternalServerError()
		log.Printf("Account hasn't been created! err: %s", err)
		return
	}

	idJSON, err := json.Marshal(accountID)
	if err != nil {
		resp.SetStatusInternalServerError()
		log.Println("Data hasn't been encoded to JSON!")
		return
	}

	resp.SetStatusOk()
	resp.SetData(idJSON)
}
