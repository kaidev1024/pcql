package pcql

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx/v3"
)

type UUID = gocql.UUID

type Batch = *gocqlx.Batch

type M = qb.M
