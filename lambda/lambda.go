package lambda

import (
  "encoding/json"
  "errors"

  "github.com/queeno/dob-api/app"
  "github.com/queeno/dob-api/db"

  "github.com/aws/aws-lambda-go/events"
)

type lambdaInterface interface {
  HandleGetUser(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
  HandlePutUser(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

type Lambda struct {
  lambda lambdaInterface
  app app.MyApp
}

type userBirthday struct {
  DateOfBirth string    `json:"dateOfBirth"`
}

type messageResponse struct {
  Message string        `json:"message"`
}

func (l Lambda) HandleRouteRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
  if req.HTTPMethod == "PUT" {
    return l.handlePutUser(req)
  } else if req.HTTPMethod == "GET" {
    return l.handleGetUser(req)
  }
  return events.APIGatewayProxyResponse{}, errors.New("Method not supported")
}

func (l Lambda) handlePutUser(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  username := req.PathParameters["username"]

  if req.Body == "" {
    eM := "Please add a body to the request in the form: {\"dateOfBirth\": \"YYYY-MM-DD\" }"
    return internalServerError(eM)
  }

  var uB userBirthday

  err := json.Unmarshal([]byte(req.Body), &uB)
  if err != nil {
    return internalServerError(err.Error())
  }

  err = l.app.UpdateUsername(username, uB.DateOfBirth)
  if err != nil {
    return internalServerError(err.Error())
  }

  return clientNoContent()
}

func (l Lambda) handleGetUser(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  username := req.PathParameters["username"]

  message, err := l.app.GetDateOfBirth(username)
  if err != nil {
    return internalServerError(err.Error())
  }

  mR := &messageResponse{
    Message: message,
  }

  jsonResponse, err := json.Marshal(mR)
  if err != nil {
    return internalServerError(err.Error())
  }

  return clientStatusOK(string(jsonResponse))
}

func NewLambda() *Lambda {
  return &Lambda{
    app: app.NewApp(db.NewDynamoDB()),
  }
}

func clientStatusOK(message string) (events.APIGatewayProxyResponse, error) {
  return events.APIGatewayProxyResponse{
    StatusCode: 200,
    Body: message,
  }, nil
}

func clientNoContent() (events.APIGatewayProxyResponse, error) {
  return events.APIGatewayProxyResponse{
    StatusCode: 204,
  }, nil
}

func internalServerError(message string) (events.APIGatewayProxyResponse, error) {
  return events.APIGatewayProxyResponse{
    StatusCode: 500,
    Body:       message,
  }, nil
}
