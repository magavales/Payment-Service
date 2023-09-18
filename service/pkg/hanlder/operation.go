package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"service/pkg/database"
	"service/pkg/model"
	"service/pkg/response"
	"strconv"
)

func (h *Handler) getBalance(ctx *gin.Context) {
	var (
		account model.Account
		resp    response.Response
		db      database.Database
		err     error
	)

	resp.RespWriter = ctx.Writer

	userID, _ := strconv.Atoi(ctx.Param("user_id"))

	db.Connect()
	account, err = db.Access.GetData(db.Pool, int64(userID))
	if err != nil {
		log.Printf("Table doesn't have rows with id = %d", userID)
		resp.SetStatusNotFound()
		return
	}

	accountToJSON, err := json.Marshal(account)
	if err != nil {
		resp.SetStatusNotFound()
		log.Println("Data hasn't been encoded to JSON!")
		return
	}

	resp.SetStatusOk()
	resp.SetData(accountToJSON)
}

func (h *Handler) createAccount(ctx *gin.Context) {
	var (
		data model.DataRequest
		resp response.Response
		db   database.Database
		err  error
	)

	err = data.DecodeJSON(ctx.Request.Body)
	if err != nil {
		log.Println("JSON hasn't been decoded!")
		resp.SetStatusBadRequest()
		return
	}

	resp.RespWriter = ctx.Writer
	db.Connect()

	err = db.Access.AddData(db.Pool, data.UserId, data.Amount)
	if err != nil {
		resp.SetStatusBadRequest()
		log.Fatalf("Can't append data to table! err: %s.", err)
		return
	}
	log.Printf("Add data to table!")
	resp.SetStatusOk()
}

func (h *Handler) updateBalance(ctx *gin.Context) {
	var (
		data    model.DataRequest
		db      database.Database
		account model.Account
		resp    response.Response
		err     error
	)

	err = data.DecodeJSON(ctx.Request.Body)
	if err != nil {
		log.Println("JSON hasn't been decoded!")
		resp.SetStatusBadRequest()
		return
	}

	resp.RespWriter = ctx.Writer
	db.Connect()

	operation := ctx.Param("operations_id")

	switch operation {
	case string(model.Increase):
		err = db.Access.IncreaseData(db.Pool, data.UserId, data.Amount)
		if err != nil {
			log.Printf("I can't communicate with the database. err: %s.", err)
			resp.SetStatusBadRequest()
			return
		}
		resp.SetStatusOk()
	case string(model.Decrease):
		account, err = db.Access.GetData(db.Pool, data.UserId)
		if err != nil {
			log.Printf("Table doesn't have rows with id = %d.", data.UserId)
			resp.SetStatusNotFound()
			return
		}
		if account.Balance >= data.Amount {
			err = db.Access.DecreaseData(db.Pool, data.UserId, data.Amount)
			if err != nil {
				log.Printf("I can't communicate with the database. err: %s", err)
				resp.SetStatusBadRequest()
				return
			}
			resp.SetStatusOk()
			return
		} else {
			log.Printf("The balance is less than the amount requested.")
			resp.SetStatusBadRequest()
		}
	case string(model.Transfer):
		account, err = db.Access.GetData(db.Pool, data.FromID)
		if err != nil {
			log.Printf("Table doesn't have rows with id = %d.", data.UserId)
			resp.SetStatusNotFound()
			return
		}
		if account.Balance >= data.Amount {
			err = db.Access.DecreaseData(db.Pool, data.FromID, data.Amount)
			if err != nil {
				log.Printf("I can't communicate with database. err: %s", err)
				resp.SetStatusBadRequest()
				return
			} else {
				err = db.Access.IncreaseData(db.Pool, data.ToID, data.Amount)
				if err != nil {
					log.Printf("I can't communicate with the database. err: %s", err)
					resp.SetStatusBadRequest()
					return
				}
			}
			log.Printf("Transfer has been completed")
			resp.SetStatusOk()
			return
		} else {
			log.Printf("The balance is less than the amount requested.")
			resp.SetStatusBadRequest()
		}
	}
}
