package pcql

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
)

func CreateNewBatch() *gocqlx.Batch {
	return sessionx.NewBatch(gocql.LoggedBatch)
}

func ExecuteBatch(batch *gocqlx.Batch) error {
	return sessionx.ExecuteBatch(batch)
}
