package pcql

import (
	"fmt"
	"time"

	astra "github.com/datastax/gocql-astra"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
)

var placesSessionx gocqlx.Session
var placesSession *gocql.Session

func SetupCassandraPlaces(id, token, keyspace string) error {
	cluster, err := astra.NewClusterFromURL("https://api.astra.datastax.com", id, token, 10*time.Second)
	if err != nil {
		return fmt.Errorf("unable to create cassandra(astra) cluster: %v", err)
	}

	cluster.Keyspace = keyspace
	cluster.Timeout = 5 * time.Second
	cluster.ConnectTimeout = 5 * time.Second

	placesSession, err = cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("could not connect to astra db: %v", err)
	}
	placesSessionx, err = gocqlx.WrapSession(placesSession, err)
	if err != nil {
		return fmt.Errorf("could not wrap session: %v", err)
	}
	return nil
}

func InsertPlaces(stmt string, names []string, row any) error {
	return placesSessionx.Query(stmt, names).BindStruct(row).ExecRelease()
}

func GetPlaces(stmt string, names []string, input, ret any) error {
	return placesSessionx.Query(stmt, names).BindStruct(input).GetRelease(ret)
}

func SelectPlaces(stmt string, names []string, input M, rows any) error {
	return placesSessionx.Query(stmt, names).BindMap(input).SelectRelease(rows)
}

func UpdatePlaces(stmt string, names []string, row any) error {
	return placesSessionx.Query(stmt, names).BindStruct(row).ExecRelease()
}
