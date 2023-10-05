package model

import (
	"encoding/json"
	"io"
)

type DataTransferRequest struct {
	FromID int64 `json:"from_id"`
	ToID   int64 `json:"to_id"`
	Amount int64 `json:"amount"`
}

func (rt *DataTransferRequest) DecodeJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&rt)

	return err
}
