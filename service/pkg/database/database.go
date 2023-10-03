package database

import (
	"github.com/jmoiron/sqlx"
	"service/pkg/database/tables"
	"service/pkg/model"
)

type Account interface {
	CreateAccount() error
	GetAccount(id int64) (model.Account, error)
	IncreaseAccountBalance(id int64, amount int64) error
	DecreaseAccountBalance(id int64, amount int64) error
}

type History interface {
	GetAllHistoryOfTransaction() ([]model.HistoryOperation, error)
}

type Database struct {
	Account
	History
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		Account: tables.NewConnectToTableAccounts(db),
		History: tables.NewConnectToTableHistory(db),
	}
}
