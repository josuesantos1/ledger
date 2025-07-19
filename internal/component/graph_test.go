package component_test

import (
	"testing"

	"github.com/josuesantos1/ledger/internal/component"
	"github.com/stretchr/testify/assert"
)

func TestGraphConnect(t *testing.T) {
	assert := assert.New(t)
	driver := component.GraphConnect("neo4j://localhost:7687", "neo4j", "password")

	assert.NotNil(driver)
}
