package service

import "service/pkg/database"

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
