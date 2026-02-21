package pcql

import (
	"fmt"
	"time"

	astra "github.com/datastax/gocql-astra"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
)

var pattriSessionx gocqlx.Session
var pattriSession *gocql.Session

func SetupCassandraPattri(id, token, keyspace string) error {
	cluster, err := astra.NewClusterFromURL("https://api.astra.datastax.com", id, token, 10*time.Second)
	if err != nil {
		return fmt.Errorf("unable to create cassandra(astra) cluster: %v", err)
	}

	cluster.Keyspace = keyspace
	cluster.Timeout = 5 * time.Second
	cluster.ConnectTimeout = 5 * time.Second

	pattriSession, err = cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("could not connect to astra db: %v", err)
	}
	pattriSessionx, err = gocqlx.WrapSession(pattriSession, err)
	if err != nil {
		return fmt.Errorf("could not wrap session: %v", err)
	}
	return nil
}

func InsertPattri(stmt string, names []string, row any) error {
	err := pattriSessionx.Query(stmt, names).BindStruct(row).ExecRelease()
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func GetPattri(stmt string, names []string, input, ret any) error {
	return pattriSessionx.Query(stmt, names).BindStruct(input).GetRelease(ret)
}

func UpdatePattri(stmt string, names []string, row any) error {
	err := pattriSessionx.Query(stmt, names).BindStruct(row).ExecRelease()
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}
