package model

type Account struct {
	User_ID int64 `json:"user_id"`
	Balance int64 `json:"balance"`
}

func (sd *Account) ParseData(values []interface{}) {
	sd.User_ID = values[0].(int64)
	sd.Balance = values[1].(int64)
}
