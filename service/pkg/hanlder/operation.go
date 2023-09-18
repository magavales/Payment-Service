package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"service/pkg/database"
	"service/pkg/model"
	"strconv"
)

func (h *Handler) getBalance(ctx *gin.Context) {
	var (
		account model.Account
		db      database.Database
		err     error
	)

	userID, _ := strconv.Atoi(ctx.Param("user_id"))

	db.Connect()
	account, err = db.Access.GetData(db.Pool, int64(userID))
	if err != nil {
		log.Printf("Table doesn't have rows with id = %d", userID)
	}

	fmt.Printf("id: %d\nbalance: %d", account.UserId, account.Balance)
}

func (h *Handler) createAccount(ctx *gin.Context) {
	var (
		data model.DataRequest
		db   database.Database
		err  error
	)

	err = data.DecodeJSON(ctx.Request.Body)
	if err != nil {
		log.Println("JSON hasn't been decoded!")
	}

	db.Connect()

	err = db.Access.AddData(db.Pool, data.UserId, data.Amount)
	if err != nil {
		log.Printf("Can't append data to table! err: %s.", err)
	}
	log.Printf("Add data to table!")
}

func (h *Handler) updateBalance(ctx *gin.Context) {
	var (
		data    model.DataRequest
		db      database.Database
		account model.Account
		err     error
	)

	err = data.DecodeJSON(ctx.Request.Body)
	if err != nil {
		log.Println("JSON hasn't been decoded!")
	}

	db.Connect()

	operation := ctx.Param("operations_id")

	switch operation {
	case string(model.Increase):
		err = db.Access.IncreaseData(db.Pool, data.UserId, data.Amount)
		if err != nil {
			log.Printf("I can't communicate with the database. err: %s.", err)
		}
	case string(model.Decrease):
		account, err = db.Access.GetData(db.Pool, data.UserId)
		if err != nil {
			log.Printf("Table doesn't have rows with id = %d.", data.UserId)
			return
		}
		if account.Balance >= data.Amount {
			err = db.Access.DecreaseData(db.Pool, data.UserId, data.Amount)
			if err != nil {
				log.Printf("I can't communicate with the database. err: %s", err)
			}
		} else {
			log.Printf("The balance is less than the amount requested.")
		}
	case string(model.Transfer):
		account, err = db.Access.GetData(db.Pool, data.FromID)
		if err != nil {
			log.Printf("Table doesn't have rows with id = %d.", data.UserId)
			return
		}
		if account.Balance >= data.Amount {
			err = db.Access.DecreaseData(db.Pool, data.FromID, data.Amount)
			if err != nil {
				log.Printf("I can't communicate with database. err: %s", err)
			}

			err = db.Access.IncreaseData(db.Pool, data.ToID, data.Amount)
			if err != nil {
				log.Printf("I can't communicate with the database. err: %s", err)
			}
		} else {
			log.Printf("The balance is less than the amount requested.")
		}
	}
}
