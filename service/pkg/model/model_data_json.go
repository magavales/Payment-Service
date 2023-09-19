package model

import (
	"encoding/json"
	"io"
)

type DataRequest struct {
	UserId string `json:"user_id"`
	FromID string `json:"from_id"`
	ToID   string `json:"to_id"`
	Amount int64  `json:"amount"`
}

func (dr *DataRequest) DecodeJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&dr)

	return err
}
