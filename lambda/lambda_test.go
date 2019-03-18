package lambda

import (
  "testing"
  "net/http"
  "errors"
  "fmt"

  "github.com/queeno/dob-api/app"

  "github.com/aws/aws-lambda-go/events"
  "github.com/stretchr/testify/assert"
)

func TestHandleRouteRequest(t *testing.T) {
  request := &events.APIGatewayProxyRequest{
    HTTPMethod: "POST",
  }

  lambda := &Lambda{}

  _, err := lambda.HandleRouteRequest(*request)

  fmt.Println("Asserting that POST results in error")
  assert.Error(t, err)
}

func TestHandlePutUser(t *testing.T) {
  tcs := []struct{
      Username string
      Dob string
      Status int
      Body string
  }{
      {"simon", "2001-01-01", http.StatusNoContent, ""},
      {"karl", "abcd-ef-gh", http.StatusInternalServerError, "The date of birth provided abcd-ef-gh didn't validate"},
      {"j0sh", "2011-01-01", http.StatusInternalServerError, "The username provided j0sh didn't validate"},
      {"", "2011-01-01", http.StatusInternalServerError, "The username provided  didn't validate"},
  }

  app := &app.MockMyApp{}

  app.
    On("UpdateUsername", "simon", "2001-01-01").
    Return(nil).
    On("UpdateUsername", "karl", "abcd-ef-gh").
    Return(errors.New("The date of birth provided abcd-ef-gh didn't validate")).
    On("UpdateUsername", "j0sh", "2011-01-01").
    Return(errors.New("The username provided j0sh didn't validate")).
    On("UpdateUsername", "", "2011-01-01").
    Return(errors.New("The username provided  didn't validate"))

  lambda := &Lambda{
    app: app,
  }

  for _, tc := range(tcs) {
    request := &events.APIGatewayProxyRequest{
      Body: fmt.Sprintf("{ \"dateOfBirth\": \"%s\" }", tc.Dob),
      HTTPMethod: "PUT",
      PathParameters: map[string]string{
        "username": tc.Username,
      },
    }

    response, err := lambda.HandleRouteRequest(*request)

    fmt.Println(fmt.Sprintf("Expected Status: %d. Got: %d", tc.Status, response.StatusCode))
    assert.Equal(t, tc.Status, response.StatusCode)

    if response.Body != "" {
      fmt.Println(fmt.Sprintf("Expected Body: %s. Got: %s", tc.Body, response.Body))
      assert.Equal(t, tc.Body, response.Body)
      assert.NoError(t, err)
    } else {
      fmt.Println(fmt.Sprintf("Expected Body: %s. Got: %s", tc.Body, response.Body))
      assert.Equal(t, tc.Body, response.Body)
    }
  }
}

func TestHandleGetUser(t *testing.T) {
  tcs := []struct{
      Username string
      Status int
      Body string
  }{
    {"simon", http.StatusOK, "{\"message\":\"Hello simon! Your birthday is in 10 day(s)\"}"},
    {"karl", http.StatusInternalServerError, "Couldn't find username"},
    {"j0sh", http.StatusInternalServerError, "Couldn't find username"},
    {"", http.StatusInternalServerError, "Couldn't find username"},
  }

  app := &app.MockMyApp{}

  app.
    On("GetDateOfBirth", "simon").
    Return("Hello simon! Your birthday is in 10 day(s)", nil).
    On("GetDateOfBirth", "karl").
    Return("", errors.New("Couldn't find username")).
    On("GetDateOfBirth", "j0sh").
    Return("", errors.New("Couldn't find username")).
    On("GetDateOfBirth", "").
    Return("", errors.New("Couldn't find username"))


  lambda := &Lambda{
    app: app,
  }

  for _, tc := range(tcs) {
    request := &events.APIGatewayProxyRequest{
      Body: tc.Body,
      HTTPMethod: "GET",
      PathParameters: map[string]string{
        "username": tc.Username,
      },
    }

    response, err := lambda.HandleRouteRequest(*request)

    fmt.Println(fmt.Sprintf("Expected Status: %d. Got: %d", tc.Status, response.StatusCode))
    assert.Equal(t, tc.Status, response.StatusCode)

    if response.Body != "" {
      fmt.Println(fmt.Sprintf("Expected Body: %s. Got: %s", tc.Body, response.Body))
      assert.Equal(t, tc.Body, response.Body)
      assert.NoError(t, err)
    } else {
      fmt.Println(fmt.Sprintf("Expected Body: %s. Got: %s", tc.Body, response.Body))
      assert.Nil(t, tc.Body, response.Body)
    }
  }
}
