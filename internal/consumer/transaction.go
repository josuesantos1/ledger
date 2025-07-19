package consumer

import (
	"fmt"
	"log"

	"github.com/josuesantos1/ledger/internal/component"
)

type Transaction struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
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

			transaction := Transaction{
				ID:     string(d.Body),
				Amount: 100.0,
			}

			fmt.Printf("Processing transaction ID: %s with amount: %.2f\n", transaction.ID, transaction.Amount)

		}
	}()
}
