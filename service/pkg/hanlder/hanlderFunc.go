package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"service/pkg/database"
	"service/pkg/model"
)

func (h *Handler) Pay(ctx *gin.Context) {
	var (
		data    model.DataRequest
		account model.Account
		db      database.Database
		err     error
	)

	err = data.DecodeJSON(ctx.Request.Body)
	if err != nil {
		log.Println("JSON hasn't been decoded!")
	}
	db.Connect()

	switch data.Operation {
	case string(model.Increase):
		err = db.Access.IncreaseData(db.Pool, data)
		if err != nil {
			log.Printf("I can't communicate with database. err: %s", err)
		}
	case string(model.Decrease):
		account, err = db.Access.GetData(db.Pool, data)
		if err != nil {
			log.Printf("I can't communicate with database. err: %s", err)
		}
		if account.Balance > data.Amount {
			err = db.Access.DecreaseData(db.Pool, data)
			if err != nil {
				log.Printf("I can't communicate with database. err: %s", err)
			}
		} else {
			log.Printf("The balance is less than the requested amount.")
		}
	}
}
