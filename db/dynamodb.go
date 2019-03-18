package db

import (
  "errors"

  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/aws"
)

type DynamoDB struct {
  db dynamodbiface.DynamoDBAPI
  Session *session.Session
}

func (d DynamoDB) Get(key string) (string, error) {

  q := &dynamodb.GetItemInput{
    TableName: aws.String("DateOfBirths"),
    Key: map[string]*dynamodb.AttributeValue{
      "username": {
        S: aws.String(key),
      },
    },
  }

  result, err := d.db.GetItem(q)
  if err != nil {
    return "", err
  }

  dob := result.Item["dateOfBirth"]
  if dob == nil {
    return "", errors.New("dateOfBirth not set in DB")
  }

  return *dob.S, nil
}

func (d DynamoDB) Put(key string, value string) error {

  input := &dynamodb.PutItemInput{
    TableName:  aws.String("DateOfBirths"),
    Item:       map[string]*dynamodb.AttributeValue{
			"dateOfBirth": {
				S: aws.String(value),
			},
      "username": {
        S: aws.String(key),
      },
		},
  }

  _, err := d.db.PutItem(input)
  if err != nil {
    return err
  }
  return nil
}

func (d DynamoDB) Close() {}

func (d *DynamoDB) Open() error {
  dynamoDBsvc := dynamodb.New(d.Session)
  db := dynamodbiface.DynamoDBAPI(dynamoDBsvc)

  d.db = db

  return nil
}

func NewDynamoDB() *DynamoDB {
  db := &DynamoDB{
    Session: session.Must(session.NewSession(&aws.Config{
    	Region: aws.String("eu-west-2"),
    })),
  }

  _ = db.Open()
  return db
}
