package pcql

import (
	"errors"

	"github.com/gocql/gocql"
)

func IsErrorNotFound(err error) bool {
	return errors.Is(err, gocql.ErrNotFound)
}
