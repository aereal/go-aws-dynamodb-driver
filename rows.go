package ddbdriver

import (
	"database/sql/driver"
	"reflect"
)

type ddbRows struct{}

var _ interface {
	driver.Rows
	driver.RowsColumnTypeDatabaseTypeName
	driver.RowsColumnTypeScanType
} = &ddbRows{}

func (r *ddbRows) Columns() []string {
	panic("not implemented") // TODO: Implement
}

func (r *ddbRows) ColumnTypeDatabaseTypeName(index int) string {
	panic("not implemented") // TODO: Implement
}

func (r *ddbRows) ColumnTypeScanType(index int) reflect.Type {
	panic("not implemented") // TODO: Implement
}

func (r *ddbRows) Close() error {
	return nil
}

func (r *ddbRows) Next(dest []driver.Value) error {
	panic("not implemented") // TODO: Implement
}
