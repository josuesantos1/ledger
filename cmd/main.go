package main

import (
	"context"

	"github.com/josuesantos1/ledger/internal/component"
)

func main() {
	ctx := context.Background()
	driver := component.GraphConnect("neo4j://localhost:7687", "neo4j", "pwd12345")

	defer driver.Close(ctx)

}
