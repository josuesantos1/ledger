package component

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GraphConnect(uri string, username string, password string) neo4j.DriverWithContext {
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Fatalf("Failed to create driver: %v", err)
	}
	return driver
}
