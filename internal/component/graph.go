package component

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (c *Component) GraphConnect(uri string, username string, password string) {
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		panic("Failed to create driver: " + err.Error())
	}

	c.GraphConn = driver
}
