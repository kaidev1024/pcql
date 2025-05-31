package pcql

import "github.com/scylladb/gocqlx/v3/table"

func CreateTable(name string, partKeys, sortKeys, otherCols []string) *table.Table {
	m := table.Metadata{
		Name:    name,
		PartKey: partKeys,
		SortKey: sortKeys,
		Columns: append(append(partKeys, sortKeys...), otherCols...),
	}
	return table.New(m)
}
