package consumer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/josuesantos1/ledger/internal/component"
)

type Amount struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
	Fee      float64 `json:"fee"`
}

type Transaction struct {
	ID              string `json:"id"`
	TransactionType string `json:"transaction_type"`
	DebitAmount     Amount `json:"debit_amount"`
	CreditAmount    Amount `json:"credit_amount"`
	CreatedAt       string `json:"created_at"`
	AccountID       string `json:"account_id"`
}

func (t *Transaction) Process(component *component.Component) {
	msg, err := component.QueueChan.Consume(
		"create-transaction",
		"",
		true,  // auto-acknowledge
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		log.Fatalf("Failed to consume messages from RabbitMQ: %v", err)
		return
	}

	fmt.Println("Transaction consumer started, waiting for messages...")

	go func() {
		for d := range msg {
			fmt.Printf("Received a message: %s\n", d.Body)

			var transaction Transaction
			if err := json.Unmarshal(d.Body, &transaction); err != nil {
				log.Printf("Failed to unmarshal transaction: %v", err)
				continue
			}

			fmt.Printf("Processing transaction ID: %s with debit amount: %.2f and credit amount: %.2f\n",
				transaction.ID, transaction.DebitAmount.Value, transaction.CreditAmount.Value)

		}
	}()
}
