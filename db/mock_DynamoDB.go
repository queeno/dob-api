package db

import (
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/aws"
)

type mockDynamoDBClient struct {
  dynamodbiface.DynamoDBAPI
}

func (m mockDynamoDBClient) PutItem(item *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error){
  return nil, nil
}

func (m mockDynamoDBClient) GetItem(item *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error){

  gio := &dynamodb.GetItemOutput{
    Item: map[string]*dynamodb.AttributeValue{
      "username": {
        S: item.Key["username"].S,
      },
      "dateOfBirth": {
        S: aws.String("2011-01-01"),
      },
    },
  }

  return gio, nil
}
