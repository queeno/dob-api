package api

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "fmt"
    "errors"
    "strings"
    "io/ioutil"

    "github.com/queeno/dob-api/app"

    "github.com/stretchr/testify/assert"
)

func TestPutUser(t *testing.T) {
  tcs := []struct{
      Username string
      Dob string
      Status int
      Body string
  }{
      {"simon", "2001-01-01", http.StatusNoContent, ""},
      {"karl", "abcd-ef-gh", http.StatusInternalServerError, "The date of birth provided abcd-ef-gh didn't validate"},
      {"j0sh", "2011-01-01", http.StatusInternalServerError, "The username provided j0sh didn't validate"},
      {"", "2011-01-01", http.StatusNotFound, "404 page not found\n"},
  }

  app := &app.MockMyApp{}

  app.
    On("UpdateUsername", "simon", "2001-01-01").
    Return(nil).
    On("UpdateUsername", "karl", "abcd-ef-gh").
    Return(errors.New("The date of birth provided abcd-ef-gh didn't validate")).
    On("UpdateUsername", "j0sh", "2011-01-01").
    Return(errors.New("The username provided j0sh didn't validate"))

  api := &Api{
    router: newRouter(),
    app: app,
  }

  api.addRoutes()

  for _, tc := range(tcs) {
    path := fmt.Sprintf("/hello/%s", tc.Username)
    req, err := http.NewRequest("PUT", path, strings.NewReader(fmt.Sprintf("{ \"dateOfBirth\": \"%s\" }", tc.Dob)))
    if err != nil {
        t.Fatal(err)
    }

    resp := httptest.NewRecorder()
    api.router.ServeHTTP(resp, req)

    fmt.Println(fmt.Sprintf("Expected Status: %d. Got: %d", tc.Status, resp.Code))
    assert.Equal(t, tc.Status, resp.Code)
    body := resp.Body

    if body != nil {
      bodyContent, err := ioutil.ReadAll(body)
      if err != nil {
        t.Fatal(err)
      }
      fmt.Println(fmt.Sprintf("Expected Body: %s. Got: %s", tc.Body, bodyContent))
      assert.Equal(t, tc.Body, string(bodyContent))
    } else {
      fmt.Println(fmt.Sprintf("Expected Body: %s. Got: %s", tc.Body, body))
      assert.Nil(t, tc.Body, body)
    }
  }
}

func TestGetUser(t *testing.T) {
  tcs := []struct{
      Username string
      Status int
      Body string
  }{
    {"simon", http.StatusOK, "{\"message\":\"Hello simon! Your birthday is in 10 day(s)\"}"},
    {"karl", http.StatusInternalServerError, "Couldn't find username"},
    {"j0sh", http.StatusInternalServerError, "Couldn't find username"},
    {"", http.StatusNotFound, "404 page not found\n"},
  }

  app := &app.MockMyApp{}

  app.
    On("GetDateOfBirth", "simon").
    Return("Hello simon! Your birthday is in 10 day(s)", nil).
    On("GetDateOfBirth", "karl").
    Return("", errors.New("Couldn't find username")).
    On("GetDateOfBirth", "j0sh").
    Return("", errors.New("Couldn't find username"))

  api := &Api{
    router: newRouter(),
    app: app,
  }

  api.addRoutes()

  for _, tc := range(tcs) {
    path := fmt.Sprintf("/hello/%s", tc.Username)
    req, err := http.NewRequest("GET", path, nil)
    if err != nil {
        t.Fatal(err)
    }

    resp := httptest.NewRecorder()
    api.router.ServeHTTP(resp, req)

    fmt.Println(fmt.Sprintf("Expected Status: %d. Got: %d", tc.Status, resp.Code))
    assert.Equal(t, tc.Status, resp.Code)
    body := resp.Body

    if body != nil {
      bodyContent, err := ioutil.ReadAll(body)
      if err != nil {
        t.Fatal(err)
      }
      fmt.Println(fmt.Sprintf("Expected Body: %s. Got: %s", tc.Body, string(bodyContent)))
      assert.Equal(t, tc.Body, string(bodyContent))
    } else {
      fmt.Println(fmt.Sprintf("Expected Body: %s. Got: %s", tc.Body, body))
      assert.Nil(t, tc.Body, body)
    }
  }
}
