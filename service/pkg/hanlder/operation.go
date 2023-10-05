package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"service/pkg/model"
	"service/pkg/response"
	"strconv"
)

func (h *Handler) getHistory(ctx *gin.Context) {
	var (
		historyTransactions []model.HistoryTransaction
		historyToJSON       []byte
		resp                response.Response
		err                 error
	)

	str := ctx.Param("user_id")
	id, _ := strconv.Atoi(str)
	resp.RespWriter = ctx.Writer

	historyTransactions, err = h.service.Payment.GetAllHistoryOfTransaction(int64(id))
	if err != nil {
		resp.SetStatusInternalServerError()
		log.Printf("Informantion about account hasn't been gotten. err: %s", err)
		return
	}
	if len(historyTransactions) == 0 {
		resp.SetStatusBadRequest()
		log.Printf("Database don't have history transaction of account with id=%d. err: %s", id, err)
		return
	}

	historyToJSON, err = json.Marshal(historyTransactions)
	if err != nil {
		resp.SetStatusInternalServerError()
		log.Println("Data hasn't been encoded to JSON!")
		return
	}

	resp.SetStatusOk()
	resp.SetData(historyToJSON)
}
