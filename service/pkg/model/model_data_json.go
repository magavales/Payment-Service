package model

import (
	"encoding/json"
	"io"
)

type DataRequest struct {
	User_ID   int64  `json:"user_id"`
	Amount    int64  `json:"amount"`
	Operation string `json:"operation"`
}

func (dr *DataRequest) DecodeJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&dr)

	return err
}
