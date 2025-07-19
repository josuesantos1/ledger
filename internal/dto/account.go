package dto

type Account struct {
	ID        string `json:"id"`
	TaxId     string `json:"tax_id"`
	Currency  string `json:"currency"`
	Country   string `json:"country"`
	CreatedAt string `json:"created_at"`
}
