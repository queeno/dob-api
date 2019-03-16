package app

import (
  "testing"
  "fmt"

  "github.com/stretchr/testify/assert"
)

func TestDateOfBirth(t *testing.T) {
  tcs := []struct{
      Dob string
      Match bool
      Error error
  }{
      {"2012-1-01", false, nil},
      {"2012-01-", false, nil},
      {"201-01-1", false, nil},
      {"a012-01-01", false, nil},
      {"2012-01-011", false, nil},
      {"2012-01-01a", false, nil},
      {"20122-01-01", false, nil},
      {"2012-011-01", false, nil},
      {"2012/01/01", false, nil},
      {"2012-01-01", true, nil},
      {"2099-01-01", false, nil},
      {"", false, nil},
  }

  for _, tc := range tcs {
    dbv := new(DobValidator)

    fmt.Println(fmt.Sprintf("Testing dob: %s",tc.Dob))
    match, err := dbv.validateDateOfBirth(tc.Dob)

    if tc.Error != nil {
      assert.Error(t, err, tc.Error)
    } else {
      assert.Equal(t, match, tc.Match)
      fmt.Println(fmt.Sprintf("PASS: Result match: %t", match))
    }
  }
}


func TestUsername(t *testing.T) {
  tcs := []struct{
      Username string
      Match bool
      Error error
  }{
      {"simon", true, nil},
      {"SIMon", true, nil},
      {"SIMON", true, nil},
      {"s1m0n", false, nil},
      {"S1m0n", false, nil},
      {"", false, nil},
  }

  for _, tc := range tcs {
    dbv := new(DobValidator)

    fmt.Println(fmt.Sprintf("Testing username: %s, match: %t",tc.Username, tc.Match))
    match, err := dbv.validateUsername(tc.Username)

    if tc.Error != nil {
      assert.Error(t, err, tc.Error)
    } else {
      assert.Equal(t, match, tc.Match)
      fmt.Println(fmt.Sprintf("PASS: Result match: %t", match))
    }
  }
}
