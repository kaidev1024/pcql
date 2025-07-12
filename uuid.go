package pcql

import (
	"bytes"

	"github.com/gocql/gocql"
)

var ZeroUUID UUID = gocql.UUID{}

func IsZeroUUID(u gocql.UUID) bool {
	return u == gocql.UUID{}
}

func TimeUUID() UUID {
	return gocql.TimeUUID()
}

func IsUUIDGreater(uuid1, uuid2 gocql.UUID) bool {
	return bytes.Compare(uuid1[:], uuid2[:]) > 0
}
