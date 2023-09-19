package model

type Account struct {
	UserId  string `json:"user_id"`
	Balance int64  `json:"balance"`
}

func (sd *Account) ParseData(values []interface{}) {
	sd.UserId = values[0].(string)
	sd.Balance = values[1].(int64)
}
