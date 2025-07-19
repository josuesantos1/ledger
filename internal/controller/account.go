package controller

import (
	"log"

	"github.com/josuesantos1/ledger/internal/component"
	"github.com/josuesantos1/ledger/internal/dto"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func ProcessAccount(component *component.Component, account *dto.Account) {
	if account == nil {
		log.Println("Received nil account, skipping processing")
		return
	}

	result, err := neo4j.ExecuteQuery(component.Ctx, component.GraphConn,
		`CREATE (a:Account {id: $id, tax_id: $tax_id, currency: $currency, country: $country, created_at: $created_at}) RETURN a`,
		map[string]any{
			"id":         account.ID,
			"tax_id":     account.TaxId,
			"currency":   account.Currency,
			"country":    account.Country,
			"created_at": account.CreatedAt,
		},
		neo4j.EagerResultTransformer,
	)

	if err != nil {
		log.Printf("Failed to create account in Neo4j: %v\n", err)
		return
	}

	log.Printf("Account created successfully: %v\n", result)
}
