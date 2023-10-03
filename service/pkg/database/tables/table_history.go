package tables

import (
	"github.com/jmoiron/sqlx"
	"service/pkg/model"
)

type TableHistory struct {
	db *sqlx.DB
}

func NewConnectToTableHistory(db *sqlx.DB) *TableHistory {
	return &TableHistory{db: db}
}

func (th *TableHistory) GetAllHistoryOfTransaction() ([]model.HistoryOperation, error) {
	var history []model.HistoryOperation

	err := th.db.Get(&history, "SELECT * FROM history")
	if err != nil {
		return nil, err
	}

	return history, nil
}
