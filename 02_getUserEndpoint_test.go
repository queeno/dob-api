package main

import (
  "testing"
  "io/ioutil"
  "net/http"
  "fmt"

  "github.com/stretchr/testify/assert"
)

func TestGetUserEndpoint(t *testing.T) {

  tcs := []struct{
      Username string
      Status int
      ResponseBody string
  }{
      {"michael", http.StatusOK, "{\"message\":\"Hello, michael! Happy birthday!\"}"},
      {"karl", http.StatusInternalServerError, "User karl doesn't exist in the database"},
      {"j0sh", http.StatusInternalServerError, "The username provided j0sh didn't validate"},
      {"", http.StatusNotFound, "404 page not found\n"},
  }

  for _, tc := range tcs {
    client := new(http.Client)
    path := fmt.Sprintf("http://localhost:8000/hello/%s", tc.Username)

    req, err := http.NewRequest("GET", path, nil)
    if err != nil {
      t.Fatal(err)
    }

    resp, err := client.Do(req)
    if err != nil {
      t.Fatal(err)
    }

    fmt.Println(fmt.Sprintf("Expected Status: %d. Got: %d", tc.Status, resp.StatusCode))
    assert.Equal(t, tc.Status, resp.StatusCode)

    if resp.Body != nil {
      bodyContent, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        t.Fatal(err)
      }
      if resp.StatusCode == http.StatusOK {
        assert.Equal(t, tc.ResponseBody, string(bodyContent))
      } else {
        fmt.Println(fmt.Sprintf("Expected Body: %s. Got: %s", tc.ResponseBody, bodyContent))
        assert.Equal(t, tc.ResponseBody, string(bodyContent))
      }
    } else {
      fmt.Println(fmt.Sprintf("Expected Body: %s. Got: %s", tc.ResponseBody, resp.Body))
      assert.Nil(t, tc.ResponseBody, resp.Body)
    }
  }
}
