package model

type Transaction string

const (
	Increase Transaction = "increase"
	Decrease Transaction = "decrease"
	Transfer Transaction = "transfer"
)
