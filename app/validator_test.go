package app

import (
  "testing"
  "fmt"
  "time"

  "github.com/stretchr/testify/assert"
)

func TestDateOfBirth(t *testing.T) {
  tcs := []struct{
      Dob string
      Today time.Time
      Match bool
      Error error
  }{
      {"2012-1-01", time.Date(1975, 12, 21, 0, 0, 0, 0, time.UTC), false, nil},
      {"2012-01-", time.Date(1975, 12, 21, 0, 0, 0, 0, time.UTC), false, nil},
      {"201-01-1", time.Date(1975, 12, 21, 0, 0, 0, 0, time.UTC), false, nil},
      {"a012-01-01", time.Date(1975, 12, 21, 0, 0, 0, 0, time.UTC), false, nil},
      {"2012-01-011", time.Date(1975, 12, 21, 0, 0, 0, 0, time.UTC), false, nil},
      {"2012-01-01a", time.Date(1975, 12, 21, 0, 0, 0, 0, time.UTC), false, nil},
      {"2012-01-01a", time.Date(1975, 12, 21, 0, 0, 0, 0, time.UTC), false, nil},
      {"20122-01-01", time.Date(1975, 12, 21, 0, 0, 0, 0, time.UTC), false, nil},
      {"2012-011-01", time.Date(1975, 12, 21, 0, 0, 0, 0, time.UTC), false, nil},
      {"2012/01/01", time.Date(1975, 12, 21, 0, 0, 0, 0, time.UTC), false, nil},
      {"2012-01-01", time.Date(1975, 12, 21, 0, 0, 0, 0, time.UTC), false, nil},
      {"2099-01-01", time.Date(1975, 12, 21, 0, 0, 0, 0, time.UTC), false, nil},
      {"", time.Date(1975, 12, 21, 0, 0, 0, 0, time.UTC), false, nil},
      {"2099-01-01", time.Date(2010, 12, 21, 0, 0, 0, 0, time.UTC), false, nil},
      {"1988-01-01", time.Date(2019, 12, 21, 0, 0, 0, 0, time.UTC), true, nil},
      {"0019-01-01", time.Date(2019, 12, 21, 0, 0, 0, 0, time.UTC), true, nil},
  }

  for _, tc := range tcs {
    dbv := &dobValidator{
      today: tc.Today,
    }

    fmt.Println(fmt.Sprintf("Testing dob: %s",tc.Dob))
    match, err := dbv.validateDateOfBirth(tc.Dob)

    if tc.Error != nil {
      assert.Error(t, tc.Error, err)
    } else {
      assert.Equal(t, tc.Match, match)
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
    dbv := &dobValidator{
      today: time.Now(),
    }

    fmt.Println(fmt.Sprintf("Testing username: %s, match: %t",tc.Username, tc.Match))
    match, err := dbv.validateUsername(tc.Username)

    if tc.Error != nil {
      assert.Error(t, tc.Error, err)
    } else {
      assert.Equal(t, tc.Match, match)
      fmt.Println(fmt.Sprintf("PASS: Result match: %t", match))
    }
  }
}
