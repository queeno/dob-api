package db

import (
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
  "github.com/aws/aws-sdk-go/service/dynamodb"
)

type mockDynamoDBClient struct {
  dynamodbiface.DynamoDBAPI
}

func (m mockDynamoDBClient) UpdateItem(item *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error){
  return nil, nil
}

func (m mockDynamoDBClient) GetItem(item *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error){
  i := &Item{
    Username: *item.Key["Username"].S,
    DateOfBirth: "2011-01-01",
  }

  elem, err := dynamodbattribute.MarshalMap(i)
  if err != nil {
    return nil, err
  }

  gio := &dynamodb.GetItemOutput{
    Item: elem,
  }

  return gio, nil
}
