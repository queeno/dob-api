package main

import (
  "testing"
  "io/ioutil"
  "net/http"
  "fmt"
  "strings"

  "github.com/stretchr/testify/assert"
)

func TestPutUserEndpoint(t *testing.T) {
  tcs := []struct{
      Username string
      RequestBody string
      Status int
      ResponseBody string
  }{
      {"simon", "{ \"dateOfBirth\": \"1988-03-20\" }", http.StatusNoContent, ""},
      {"karl", "{ \"dateOfBirth\": \"2007-02-29\" }", http.StatusInternalServerError, "parsing time \"2007-02-29\": day out of range"},
      {"j0sh", "{ \"dateOfBirth\": \"2007-02-29\" }", http.StatusInternalServerError, "The username provided j0sh didn't validate"},
      {"", "{ \"dateOfBirth\": \"2007-02-29\" }", http.StatusNotFound, "404 page not found\n"},
  }

  for _, tc := range tcs {
    client := new(http.Client)
    path := fmt.Sprintf("http://localhost:8000/hello/%s", tc.Username)

    requestBody := strings.NewReader(tc.RequestBody)
    req, err := http.NewRequest("PUT", path, requestBody)
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
      fmt.Println(fmt.Sprintf("Expected Body: %s. Got: %s", tc.ResponseBody, bodyContent))
      assert.Equal(t, tc.ResponseBody, string(bodyContent))
    } else {
      fmt.Println(fmt.Sprintf("Expected Body: %s. Got: %s", tc.ResponseBody, resp.Body))
      assert.Nil(t, tc.ResponseBody, resp.Body)
    }
  }
}
