package ddbdriver

import (
	"database/sql/driver"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type item = map[string]*dynamodb.AttributeValue

type ddbResult struct {
	items []item
}

var _ driver.Result = &ddbResult{}

func (r *ddbResult) LastInsertId() (int64, error) {
	panic("not implemented") // TODO: Implement
}

func (r *ddbResult) RowsAffected() (int64, error) {
	panic("not implemented") // TODO: Implement
}
