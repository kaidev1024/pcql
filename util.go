package pcql

import "github.com/gocql/gocql"

func MustParseUUID(id string) gocql.UUID {
	if id == "" {
		return gocql.UUID{}
	}
	uuid, err := gocql.ParseUUID(id)
	if err != nil {
		panic(err)
	}
	return uuid
}
