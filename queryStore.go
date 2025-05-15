package pcql

import "fmt"

type (
	Cql struct {
		Stmt  string
		Names []string
	}
	CqlStore struct {
		cqls map[CqlName]*Cql
	}
	CqlName string
)

func (qs *CqlStore) SetCql(name CqlName, stmt string, names []string) {
	qs.cqls[name] = &Cql{stmt, names}
}

func (qs *CqlStore) GetCql(name CqlName) *Cql {
	cql, ok := qs.cqls[name]
	if !ok {
		panic(fmt.Sprintf("%s not found", name))
	}
	return cql
}

func CreateCqlStore() *CqlStore {
	return &CqlStore{make(map[CqlName]*Cql)}
}
