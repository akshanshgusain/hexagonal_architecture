package domain

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionId   string
	AccountId       string
	Amount          float64
	TransactionType string
	TransactionDate string
}

type TransactionRepository interface {
}