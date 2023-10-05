package model

import (
	"encoding/json"
	"io"
)

type DataRequest struct {
	UserId int64 `json:"user_id"`
	FromID int64 `json:"from_id"`
	ToID   int64 `json:"to_id"`
	Amount int64 `json:"amount"`
}

func (dr *DataRequest) DecodeJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&dr)

	return err
}
