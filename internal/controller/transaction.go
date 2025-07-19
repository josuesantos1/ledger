package controller

import (
	"log"

	"github.com/josuesantos1/ledger/internal/component"
	"github.com/josuesantos1/ledger/internal/dto"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func ProcessTransaction(component *component.Component, transaction *dto.Transaction) {
	if transaction == nil {
		log.Println("Received nil transaction, skipping processing")
		return
	}

	query := `
		MATCH (a:Account {id: $accountId})
		CREATE (t:Transaction {
			id: $id,
			debitValue: $debitValue,
			debitCurrency: $debitCurrency,
			debitFee: $debitFee,
			creditValue: $creditValue,
			creditCurrency: $creditCurrency,
			creditFee: $creditFee,
			creditConversion: $creditConversion,
			debitConversion: $debitConversion,
			transactionDate: $transactionDate,
			transactionType: $transactionType,
			transactionId: $transactionId,
			accountId: $accountId
		})
		CREATE (t)-[:BELONGS_TO]->(a)
		RETURN t
	`

	params := map[string]any{
		"id":               transaction.ID,
		"debitValue":       transaction.DebitAmount.Value,
		"debitCurrency":    transaction.DebitAmount.Currency,
		"debitFee":         transaction.DebitAmount.Fee,
		"creditValue":      transaction.CreditAmount.Value,
		"creditCurrency":   transaction.CreditAmount.Currency,
		"creditFee":        transaction.CreditAmount.Fee,
		"creditConversion": transaction.CreditAmount.ConversionRate,
		"debitConversion":  transaction.DebitAmount.ConversionRate,
		"transactionDate":  transaction.TransactionDate,
		"transactionType":  transaction.TransactionType,
		"transactionId":    transaction.TransactionId,
		"accountId":        transaction.AccountID,
	}

	result, err := neo4j.ExecuteQuery(
		component.Ctx,
		component.GraphConn,
		query,
		params,
		neo4j.EagerResultTransformer,
	)

	if err != nil {
		log.Printf("Failed to create transaction in Neo4j: %v\n", err)
		return
	}

	log.Printf("Transaction created and linked to Account[%s]: %v\n", transaction.AccountID, result)
}
