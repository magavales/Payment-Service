package model

type HistoryOperation struct {
	UserId    string `json:"user_id"`
	Operation string `json:"operation"`
	Amount    int64  `json:"amount"`
	ToID      string `json:"to_id"`
}

func (ho *HistoryOperation) ParseData(values []interface{}) {
	ho.UserId = values[0].(string)
	ho.Operation = values[1].(string)
	ho.Amount = values[2].(int64)
	ho.ToID = values[3].(string)
}
