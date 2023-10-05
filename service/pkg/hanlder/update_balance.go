package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"service/pkg/model"
	"service/pkg/response"
)

func (h *Handler) updateBalance(ctx *gin.Context) {
	var (
		dataRequest model.DataRequest
		resp        response.Response
		err         error
		syntaxError *json.UnmarshalTypeError
		updateError error
	)
	updateError = errors.New("query hasn't been completed")
	resp.RespWriter = ctx.Writer

	err = dataRequest.DecodeJSON(ctx.Request.Body)
	if err != nil {
		if err != nil {
			if errors.As(err, &syntaxError) {
				log.Printf("JSON file has syntax error. Error: %s\n", err)
				resp.SetStatusBadRequest()
				return
			} else {
				log.Printf("The service couldn't decode JSON file. Error: %s\n", err)
				resp.SetStatusInternalServerError()
				return
			}
		}
	}

	transaction := ctx.Param("transaction_id")

	switch transaction {
	case string(model.Increase):
		err = h.service.Payment.IncreaseAccountBalance(dataRequest.UserId, dataRequest.Amount)
		if err != nil {
			if errors.As(err, &updateError) {
				log.Printf("Account with id = %d hasn't been found in database. Error: %s\n", dataRequest.UserId, err)
				resp.SetStatusNotFound()
				return
			} else {
				log.Printf("Error: %s\n", err)
				resp.SetStatusInternalServerError()
				return
			}
		} else {
			err = h.service.Payment.AddHistoryOfTransaction(dataRequest.UserId, string(model.Increase), dataRequest)
			if err != nil {
				_ = h.service.Payment.DecreaseAccountBalance(dataRequest.UserId, dataRequest.Amount)
				log.Printf("Error: %s\n", err)
				resp.SetStatusInternalServerError()
				return
			} else {
				resp.SetStatusOk()
			}
		}
	case string(model.Decrease):
		err = h.service.Payment.DecreaseAccountBalance(dataRequest.UserId, dataRequest.Amount)
		if err != nil {
			if errors.As(err, &updateError) {
				log.Printf("Account with id = %d hasn't been found in database. Error: %s\n", dataRequest.UserId, err)
				resp.SetStatusNotFound()
				return
			} else {
				log.Printf("Error: %s\n", err)
				resp.SetStatusInternalServerError()
				return
			}
		} else {
			err = h.service.Payment.AddHistoryOfTransaction(dataRequest.UserId, string(model.Decrease), dataRequest)
			if err != nil {
				_ = h.service.Payment.IncreaseAccountBalance(dataRequest.UserId, dataRequest.Amount)
				log.Printf("Error: %s\n", err)
				resp.SetStatusInternalServerError()
				return
			} else {
				resp.SetStatusOk()
			}
		}
	case string(model.Transfer):
		err = h.service.Payment.DecreaseAccountBalance(dataRequest.FromID, dataRequest.Amount)
		if err != nil {
			if errors.As(err, &updateError) {
				log.Printf("Account with id = %d hasn't been found in database. Error: %s\n", dataRequest.FromID, err)
				resp.SetStatusNotFound()
				return
			} else {
				log.Printf("Error: %s\n", err)
				resp.SetStatusInternalServerError()
				return
			}
		} else {
			err = h.service.Payment.IncreaseAccountBalance(dataRequest.ToID, dataRequest.Amount)
			if err != nil {
				_ = h.service.Payment.IncreaseAccountBalance(dataRequest.FromID, dataRequest.Amount)
				if errors.As(err, &updateError) {
					log.Printf("Account with id = %d hasn't been found in database. Error: %s\n", dataRequest.ToID, err)
					resp.SetStatusNotFound()
					return
				} else {
					log.Printf("Error: %s\n", err)
					resp.SetStatusInternalServerError()
					return
				}
			}
		}

		err = h.service.Payment.AddHistoryOfTransaction(dataRequest.FromID, string(model.Transfer), dataRequest)
		if err != nil {
			_ = h.service.Payment.IncreaseAccountBalance(dataRequest.FromID, dataRequest.Amount)
			_ = h.service.Payment.DecreaseAccountBalance(dataRequest.ToID, dataRequest.Amount)
			log.Printf("Error: %s\n", err)
			resp.SetStatusInternalServerError()
			return
		} else {
			resp.SetStatusOk()
		}
	}
}
