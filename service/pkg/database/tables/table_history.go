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

func (th *TableHistory) GetAllHistoryOfTransaction(id int64) ([]model.HistoryTransaction, error) {
	var history []model.HistoryTransaction

	err := th.db.Select(&history, "SELECT * FROM history WHERE from_id = $1", id)
	if err != nil {
		return nil, err
	}

	return history, nil
}

func (th *TableHistory) AddHistoryOfTransaction(id int64, transaction string, dataRequest model.DataRequest) error {
	_, err := th.db.Query("INSERT INTO history (from_id, to_id, transaction, amount) VALUES ($1, $2, $3, $4)",
		id, dataRequest.ToID, transaction, dataRequest.Amount)
	if err != nil {
		return err
	}

	return err
}
