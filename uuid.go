package pcql

import (
	"github.com/gocql/gocql"
)

var ZeroUUID UUID = gocql.UUID{}

func IsZeroUUID(u gocql.UUID) bool {
	return u == gocql.UUID{}
}

func TimeUUID() UUID {
	return gocql.TimeUUID()
}
