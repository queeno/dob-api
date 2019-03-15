package api

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "fmt"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
    tcs := []struct{
        routeVariable string
        returnCode int
    }{
        {"simon", http.StatusOK},
    }

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
      router := mux.NewRouter()
      router.HandleFunc("/hello/{username:[a-zA-Z0-9]+}", getUser)
      router.ServeHTTP(resp, req)

      // Asserts
      assert.Equal(t, resp.Code, http.StatusOK)
    }
}
