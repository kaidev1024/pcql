package pcql

import (
	"fmt"
	"time"

	astra "github.com/datastax/gocql-astra"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
)

var osmSessionx gocqlx.Session
var osmSession *gocql.Session

func SetupCassandraOsm(id, token, keyspace string) error {
	cluster, err := astra.NewClusterFromURL("https://api.astra.datastax.com", id, token, 10*time.Second)
	if err != nil {
		return fmt.Errorf("unable to create cassandra(astra) cluster: %v", err)
	}

	cluster.Keyspace = keyspace
	cluster.Timeout = 5 * time.Second
	cluster.ConnectTimeout = 5 * time.Second

	osmSession, err = cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("could not connect to astra db: %v", err)
	}
	osmSessionx, err = gocqlx.WrapSession(osmSession, err)
	if err != nil {
		return fmt.Errorf("could not wrap session: %v", err)
	}
	return nil
}

func InsertOsm(stmt string, names []string, row any) error {
	err := osmSessionx.Query(stmt, names).BindStruct(row).ExecRelease()
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func GetOsm(stmt string, names []string, input, ret any) error {
	return osmSessionx.Query(stmt, names).BindStruct(input).GetRelease(ret)
}

func SelectOsm(stmt string, names []string, input M, rows any) error {
	err := osmSessionx.Query(stmt, names).BindMap(input).SelectRelease(rows)
	if err != nil {
		return err
	}
	return nil
}

func UpdateOsm(stmt string, names []string, row any) error {
	err := osmSessionx.Query(stmt, names).BindStruct(row).ExecRelease()
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}
