package api

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "fmt"

    "github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
    tcs := []struct{
        routeVariable string
        returnCode int
    }{
        {"simon", http.StatusOK},
        {"SIMon", http.StatusOK},
        {"SIMON", http.StatusOK},
        {"s1m0n", http.StatusNotFound},
        {"S1m0n", http.StatusNotFound},
        {"", http.StatusNotFound},
    }

    router := newRouter()

    for _, tc := range tcs {
      // Creates the request
      path := fmt.Sprintf("/hello/%s", tc.routeVariable)
      req, err := http.NewRequest("GET", path, nil)
      if err != nil {
          t.Fatal(err)
      }

      // Records the response
      resp := httptest.NewRecorder()

      // Fires the request
      router.ServeHTTP(resp, req)

      // Asserts
      assert.Equal(t, resp.Code, tc.returnCode)
    }
}
