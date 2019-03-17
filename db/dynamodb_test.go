package db

import (
  "testing"
  "fmt"
  "net/http"
  "net/http/httptest"

  "github.com/stretchr/testify/assert"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/aws"
)

func TestOpenDB(t *testing.T) {
  var session = func() *session.Session {
  	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  		w.WriteHeader(http.StatusOK)
  	}))

  	return session.Must(session.NewSession(&aws.Config{
  		DisableSSL: aws.Bool(true),
  		Endpoint:   aws.String(server.URL),
  	}))
  }()

  db := &DynamoDB{
    Session: session,
  }

  fmt.Println("Trying to build a dynamoDB database")
  err := db.Open()

  assert.Nil(t, err)
  assert.NotNil(t, db.db)
}

func TestGetItemDynamoDB(t *testing.T) {
  db := &DynamoDB{
    db: &mockDynamoDBClient{},
  }

  fmt.Println("Getting: simon from a mocked DynamoDB")
  dob, err := db.Get("simon")
  if err != nil {
    t.Fatal(err)
  }

  assert.Equal(t, dob, "2011-01-01")
}

func TestPutItemDynamoDB(t *testing.T){
  db := &DynamoDB{
    db: &mockDynamoDBClient{},
  }

  fmt.Println("Putting: simon, 2000-11-05 into a mocked DynamoDB")
  err := db.Put("simon","2000-11-05")
  assert.NoError(t, err)
}
