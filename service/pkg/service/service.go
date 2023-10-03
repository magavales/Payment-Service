package service

import "service/pkg/database"

type Payment interface {
}

type Service struct {
	Payment
}

func NewService(db *database.Database) *Service {
	return &Service{
		Payment: NewPaymentService(db.Account, db.History),
	}
}
