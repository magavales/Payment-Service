package tables

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log"
	"service/pkg/model"
)

type TableAccounts struct {
	db *sqlx.DB
}

func NewConnectToTableAccounts(db *sqlx.DB) *TableAccounts {
	return &TableAccounts{db: db}
}

func (ta *TableAccounts) CreateAccount() error {
	_, err := ta.db.Query("INSERT INTO accounts (balance) VALUES ($1) RETURNING user_id", 0)
	if err != nil {
		return err
	}
	return err
}

func (ta *TableAccounts) GetAccount(id int64) (model.Account, error) {
	var account model.Account
	err := ta.db.Get(&account, "SELECT user_id, balance FROM accounts WHERE user_id = $1", id)
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
	}

	if account.UserId == 0 && account.Balance == 0 {
		return account, sql.ErrNoRows
	}

	return account, err
}

func (ta *TableAccounts) IncreaseAccountBalance(id int64, amount int64) error {
	tag, err := ta.db.Exec("UPDATE accounts SET balance = balance + $1 WHERE user_id = $2", amount, id)
	if err != nil {
		return err
	}

	_, err = tag.RowsAffected()
	if err != nil {
		return err
	}

	return err
}

func (ta *TableAccounts) DecreaseAccountBalance(id int64, amount int64) error {
	tag, err := ta.db.Exec("UPDATE accounts SET balance = balance - $1 WHERE user_id = $2", amount, id)
	if err != nil {
		return err
	}

	_, err = tag.RowsAffected()
	if err != nil {
		return err
	}

	return err
}
