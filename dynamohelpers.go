package lambdadynamohelpers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
)

const REGION = "REGION"

func NewSess() (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv(REGION))},
	)
	if err != nil {
		return nil, err
	}
	return dynamodb.New(sess), nil;
}

func GetItemS(keyName string, key string, table string, result interface{}) error {
	svc, err := NewSess();
	if err != nil {
		return err
	}
	av := map[string]*dynamodb.AttributeValue{}
	av[keyName] = &dynamodb.AttributeValue{
		S: aws.String(key),
	}
	output, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: av,
	})
	if err != nil {
		return err
	}
	return dynamodbattribute.UnmarshalMap(output.Item, result)
}

func PutItem(in interface{}, table string) error {
	av, err := dynamodbattribute.MarshalMap(in)
	if err != nil {
		return err
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv(REGION))},
	)
	if err != nil {
		return err
	}
	svc := dynamodb.New(sess);
	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String(table),
	}
	_, err = svc.PutItem(input)
	return err
}