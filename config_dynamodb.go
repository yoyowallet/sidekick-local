package main

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDBConfigSource struct {
	Table    string
	Key      string
	Endpoint string
	Region   string
}

func (src *DynamoDBConfigSource) List(ctx context.Context) ([]string, error) {
	var err error

	awsConfig := &aws.Config{
		Endpoint: &src.Endpoint,
		Region:   &src.Region,
	}

	sess := session.Must(session.NewSession(awsConfig))
	svc := dynamodb.New(sess)

	resp, err := svc.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		ConsistentRead: aws.Bool(true),
		Key: map[string]*dynamodb.AttributeValue{
			"key": {S: aws.String(src.Key)},
		},
		TableName: aws.String(src.Table),
	})
	if err != nil {
		return nil, err
	}

	itemsMap := make(map[string]string)
	err = dynamodbattribute.UnmarshalMap(resp.Item, &itemsMap)
	if err != nil {
		return nil, err
	}

	delete(itemsMap, src.Key)

	items := make([]string, 0, len(itemsMap))
	for k, v := range itemsMap {
		items = append(items, k+"="+v)
	}

	return items, nil
}
