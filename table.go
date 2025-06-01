package pcql

import "github.com/scylladb/gocqlx/v3/table"

func CreateTable(name string, columns, partKeys, sortKeys []string) *table.Table {
	return table.New(table.Metadata{
		Name:    name,
		Columns: columns,
		PartKey: partKeys,
		SortKey: sortKeys,
	})
}
