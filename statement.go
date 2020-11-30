package ddbdriver

import (
	"context"
	"database/sql/driver"
)

type ddbStmt struct{}

var _ interface {
	driver.Stmt
	driver.StmtExecContext
	driver.StmtQueryContext
} = &ddbStmt{}

func (s *ddbStmt) Close() error {
	return nil
}

func (s *ddbStmt) NumInput() int {
	panic("not implemented") // TODO: Implement
}

func (s *ddbStmt) Exec(args []driver.Value) (driver.Result, error) {
	nvs := make([]driver.NamedValue, len(args))
	for i, v := range args {
		nvs[i] = driver.NamedValue{Value: v}
	}
	return s.ExecContext(context.Background(), nvs)
}

func (s *ddbStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	return &ddbResult{}, nil
}

func (s *ddbStmt) Query(args []driver.Value) (driver.Rows, error) {
	nvs := make([]driver.NamedValue, len(args))
	for i, v := range args {
		nvs[i] = driver.NamedValue{Value: v}
	}
	return s.QueryContext(context.Background(), nvs)
}

func (s *ddbStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	return &ddbRows{}, nil
}
