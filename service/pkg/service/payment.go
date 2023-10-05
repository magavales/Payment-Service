package service

import (
	"service/pkg/database"
	"service/pkg/model"
)

type PaymentService struct {
	dbAccount database.Account
	dbHistory database.History
}

func NewPaymentService(dbAccount database.Account, dbHistory database.History) *PaymentService {
	return &PaymentService{
		dbAccount: dbAccount,
		dbHistory: dbHistory,
	}
}

func (ps *PaymentService) CreateAccount() (int64, error) {
	return ps.dbAccount.CreateAccount()
}

func (ps *PaymentService) GetAccount(id int64) (model.Account, error) {
	return ps.dbAccount.GetAccount(id)
}

func (ps *PaymentService) IncreaseAccountBalance(id int64, amount int64) error {
	return ps.dbAccount.IncreaseAccountBalance(id, amount)
}

func (ps *PaymentService) DecreaseAccountBalance(id int64, amount int64) error {
	return ps.dbAccount.DecreaseAccountBalance(id, amount)
}

func (ps *PaymentService) GetAllHistoryOfTransaction(id int64) ([]model.HistoryTransaction, error) {
	return ps.dbHistory.GetAllHistoryOfTransaction(id)
}

func (ps *PaymentService) AddHistoryOfTransaction(id int64, transaction string, dataRequest model.DataRequest) error {
	return ps.dbHistory.AddHistoryOfTransaction(id, transaction, dataRequest)
}
