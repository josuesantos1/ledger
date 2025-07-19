package consumer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/josuesantos1/ledger/internal/component"
	"github.com/josuesantos1/ledger/internal/controller"
	"github.com/josuesantos1/ledger/internal/dto"
)

func ConsumeAccount(component *component.Component) {
	msg, err := component.QueueChan.Consume(
		"create-account",
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

	fmt.Println("Account consumer started, waiting for messages...")

	go func() {
		for d := range msg {
			fmt.Printf("Received a message: %s\n", d.Body)

			var account *dto.Account
			if err := json.Unmarshal(d.Body, &account); err != nil {
				log.Printf("Failed to unmarshal account: %v", err)
				continue
			}

			controller.ProcessAccount(component, account)
		}
	}()
}
