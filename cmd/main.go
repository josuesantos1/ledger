package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/josuesantos1/ledger/internal/component"
	"github.com/josuesantos1/ledger/internal/consumer"
)

func main() {
	ctx := context.Background()

	component := component.Component{
		Ctx: ctx,
	}

	fmt.Println("Starting Ledger application...")

	// Connect to Neo4j
	fmt.Println("Connecting to Neo4j...")
	component.GraphConnect("neo4j://localhost:7687", "neo4j", "pwd12345")
	defer component.GraphConn.Close(ctx)
	fmt.Println("Neo4j connected successfully")

	// Connect to RabbitMQ
	fmt.Println("Connecting to RabbitMQ...")
	component.QueueConnect("amqp://guest:guest@localhost:5672/")
	component.QueueDeclare([]string{"create-transaction", "create-account"})

	defer component.QueueConn.Close()
	fmt.Println("RabbitMQ connected successfully")

	fmt.Println("All services connected. Application is running...")

	// Start the consumer
	consumer.Process(&component)
	consumer.ConsumeAccount(&component)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Printf("Application heartbeat - %s\n", time.Now().Format("15:04:05"))
			case <-ctx.Done():
				return
			}
		}
	}()

	<-sigChan
	fmt.Println("\nShutting down gracefully...")
	fmt.Println("Application stopped")
}
