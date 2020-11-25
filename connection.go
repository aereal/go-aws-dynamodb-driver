package ddbdriver

import (
	"context"
	"database/sql/driver"
)

type ddbConn struct{}

var _ interface {
	driver.Conn
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.QueryerContext
} = &ddbConn{}

func (d *ddbConn) Prepare(query string) (driver.Stmt, error) {
	return d.PrepareContext(context.Background(), query)
}

func (c *ddbConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	return &ddbStmt{}, nil
}

func (d *ddbConn) Close() error {
	return nil
}

func (d *ddbConn) Begin() (driver.Tx, error) {
	return d.BeginTx(context.Background(), driver.TxOptions{})
}

func (c *ddbConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	return &ddbTx{}, nil
}

func (c *ddbConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	return &ddbResult{}, nil
}

func (c *ddbConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	return &ddbRows{}, nil
}
