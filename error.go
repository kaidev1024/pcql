package pcql

import (
	"errors"

	"github.com/gocql/gocql"
)

func IsErrNotFound(err error) bool {
	return errors.Is(err, gocql.ErrNotFound)
}
