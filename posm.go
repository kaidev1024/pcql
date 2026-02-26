package pcql

import (
	"fmt"
	"time"

	astra "github.com/datastax/gocql-astra"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
)

var posmSessionx gocqlx.Session
var posmSession *gocql.Session

func SetupCassandraPosm(id, token, keyspace string) error {
	cluster, err := astra.NewClusterFromURL("https://api.astra.datastax.com", id, token, 10*time.Second)
	if err != nil {
		return fmt.Errorf("unable to create cassandra(astra) cluster: %v", err)
	}

	cluster.Keyspace = keyspace
	cluster.Timeout = 5 * time.Second
	cluster.ConnectTimeout = 5 * time.Second

	posmSession, err = cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("could not connect to astra db: %v", err)
	}
	posmSessionx, err = gocqlx.WrapSession(posmSession, err)
	if err != nil {
		return fmt.Errorf("could not wrap session: %v", err)
	}
	return nil
}

func InsertPosm(stmt string, names []string, row any) error {
	return posmSessionx.Query(stmt, names).BindStruct(row).ExecRelease()
}

func GetPosm(stmt string, names []string, input, ret any) error {
	return posmSessionx.Query(stmt, names).BindStruct(input).GetRelease(ret)
}

func SelectPosm(stmt string, names []string, input M, rows any) error {
	return posmSessionx.Query(stmt, names).BindMap(input).SelectRelease(rows)
}

func UpdatePosm(stmt string, names []string, row any) error {
	return posmSessionx.Query(stmt, names).BindStruct(row).ExecRelease()
}
