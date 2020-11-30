package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %+v\n", err)
		os.Exit(1)
	}
}

func run() error {
	endpoint := os.Getenv("DYNAMODB_ENDPOINT")
	if endpoint == "" {
		return errors.New("DYNAMODB_ENDPOINT is empty")
	}
	ses, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", ""),
		Endpoint:    &endpoint,
	})
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client := dynamodb.New(ses)
	switch "insert-items" {
	case "create-table":
		return createTable(ctx, client)
	case "insert-items":
		return insertItems(ctx, client)
	}
	return nil
}

func insertItems(ctx context.Context, client *dynamodb.DynamoDB) error {
	log.Printf("insertItems")
	out, err := client.ExecuteStatementWithContext(ctx, &dynamodb.ExecuteStatementInput{
		Statement: aws.String("INSERT INTO TestTable1 VALUE {'Name': 'aereal', 'Age': 30}"),
	})
	if err != nil {
		return err
	}
	fmt.Printf("output=%#v\n", out)
	return nil
}

func createTable(ctx context.Context, client *dynamodb.DynamoDB) error {
	out, err := client.CreateTableWithContext(ctx, &dynamodb.CreateTableInput{
		TableName: aws.String("TestTable1"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Name"),
				KeyType:       aws.String(dynamodb.KeyTypeHash),
			},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Name"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
	})
	if err != nil {
		return err
	}
	fmt.Printf("TableDescription=%s\n", out.TableDescription)
	return nil
}
