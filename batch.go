package pcql

import (
	"github.com/gocql/gocql"
)

func CreateNewBatch() Batch {
	return sessionx.NewBatch(gocql.LoggedBatch)
}

func ExecuteBatch(batch Batch) error {
	return sessionx.ExecuteBatch(batch)
}
