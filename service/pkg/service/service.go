package service

import (
	"service/pkg/database"
	"service/pkg/model"
)

type Payment interface {
	CreateAccount() (int64, error)
	GetAccount(id int64) (model.Account, error)
	IncreaseAccountBalance(id int64, amount int64) error
	DecreaseAccountBalance(id int64, amount int64) error
	GetAllHistoryOfTransaction(id int64) ([]model.HistoryTransaction, error)
	AddHistoryOfTransaction(id int64, transaction string, dataRequest model.DataRequest) error
}

type Service struct {
	Payment
}

func NewService(db *database.Database) *Service {
	return &Service{
		Payment: NewPaymentService(db.Account, db.History),
	}
}
