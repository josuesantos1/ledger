package dto

type Amount struct {
	Value          float64 `json:"value"`
	Currency       string  `json:"currency"`
	Fee            float64 `json:"fee"`
	ConversionRate float64 `json:"conversion_rate,omitempty"`
}

type Transaction struct {
	ID              string `json:"id"`
	TransactionType string `json:"transaction_type"`
	TransactionId   string `json:"transaction_id"`
	DebitAmount     Amount `json:"debit_amount"`
	CreditAmount    Amount `json:"credit_amount"`
	TransactionDate string `json:"transaction_date"`
	DebitAccount    string `json:"debit_account"`
	CreditAccount   string `json:"credit_account"`
}
