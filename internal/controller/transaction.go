package controller

import (
	"log"

	"github.com/josuesantos1/ledger/internal/component"
	"github.com/josuesantos1/ledger/internal/dto"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Entry struct {
	ID               string
	DebitAmount      float64
	CreditAmount     float64
	DebitCurrency    string
	CreditCurrency   string
	DebitFee         float64
	CreditFee        float64
	DebitConversion  float64
	CreditConversion float64
	TransactionDate  string
	TransactionType  string
	TransactionId    string
	AccountID        string
}

func ProcessTransaction(component *component.Component, transaction *dto.Transaction) {
	if transaction == nil {
		log.Println("Received nil transaction, skipping processing")
		return
	}

	query := `MATCH (debitAccount:Account {id: $debitAccountId}), (creditAccount:Account {id: $creditAccountId})
		CREATE (debitTransaction:Transaction {
			id: $id,
			amount: $debitAmount,
			currency: $debitCurrency,
			fee: $debitFee,
			conversion: $debitConversion,
			transactionDate: $transactionDate,
			transactionType: $transactionType,
			transactionId: $transactionId
		})
		CREATE (creditTransaction:Transaction {
			id: $id,
			amount: $creditAmount,
			currency: $creditCurrency,
			fee: $creditFee,
			conversion: $creditConversion,
			transactionDate: $transactionDate,
			transactionType: $transactionType,
			transactionId: $transactionId
		})
		RETURN debitTransaction, creditTransaction
	`

	entryDebit, entryCredit := CreateDoubleEntryTransaction(component, transaction)

	params := map[string]any{
		"debitAccountId":   entryDebit.AccountID,
		"creditAccountId":  entryCredit.AccountID,
		"id":               entryDebit.ID,
		"debitAmount":      entryDebit.DebitAmount,
		"debitCurrency":    entryDebit.DebitCurrency,
		"debitFee":         entryDebit.DebitFee,
		"debitConversion":  entryDebit.DebitConversion,
		"creditAmount":     entryCredit.CreditAmount,
		"creditCurrency":   entryCredit.CreditCurrency,
		"creditFee":        entryCredit.CreditFee,
		"creditConversion": entryCredit.CreditConversion,
		"transactionDate":  entryDebit.TransactionDate,
		"transactionType":  entryDebit.TransactionType,
		"transactionId":    entryDebit.TransactionId,
	}

	_, err := neo4j.ExecuteQuery(
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

}

func CreateDoubleEntryTransaction(component *component.Component, transaction *dto.Transaction) (Entry, Entry) {
	entryDebit := Entry{
		ID:              transaction.ID,
		DebitAmount:     transaction.DebitAmount.Value,
		DebitCurrency:   transaction.DebitAmount.Currency,
		DebitFee:        transaction.DebitAmount.Fee,
		DebitConversion: transaction.DebitAmount.ConversionRate,
		TransactionDate: transaction.TransactionDate,
		TransactionType: transaction.TransactionType,
		TransactionId:   transaction.TransactionId,
		AccountID:       transaction.DebitAccount,
	}

	entryCredit := Entry{
		ID:               transaction.ID,
		CreditAmount:     transaction.CreditAmount.Value,
		CreditCurrency:   transaction.CreditAmount.Currency,
		CreditFee:        transaction.CreditAmount.Fee,
		CreditConversion: transaction.CreditAmount.ConversionRate,
		TransactionDate:  transaction.TransactionDate,
		TransactionType:  transaction.TransactionType,
		TransactionId:    transaction.TransactionId,
		AccountID:        transaction.CreditAccount,
	}

	return entryDebit, entryCredit
}
