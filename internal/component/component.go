package component

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Component struct {
	GraphConn neo4j.DriverWithContext
	QueueConn *amqp.Connection
	QueueChan *amqp.Channel

	Ctx context.Context
}
