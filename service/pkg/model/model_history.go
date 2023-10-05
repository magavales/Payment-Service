package model

type HistoryTransaction struct {
	ID          int64  `json:"id" db:"id"`
	FromID      int64  `json:"from_id" db:"from_id"`
	ToID        int64  `json:"to_id" db:"to_id"`
	Transaction string `json:"operation" db:"transaction"`
	Amount      int64  `json:"amount" db:"amount"`
}
