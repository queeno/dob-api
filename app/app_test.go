package app

import (
  "testing"
  "fmt"
  "errors"

  "github.com/queeno/dob-api/db"

  "github.com/stretchr/testify/mock"
  "github.com/stretchr/testify/assert"
)

func TestInitialise(t *testing.T){
  db := &db.MockDatabase{}
  app := new(App)

  app.Initialise(db)

  fmt.Println("Testing app has been correctly initialised")
  assert.Equal(t, db, app.db)
  fmt.Println("PASS: App correctly initialised")
}

func TestGetDateOfBirth(t *testing.T) {

  tcs := []struct{
    Username string
    Dob string
    Error error
  }{
    {"simon", "1980-04-20", nil},
    {"anita", "", nil},
    {"s1m0n", "1980-04-20", errors.New("The username provided s1m0n didn't validate")},
  }

  validator := &MockValidator{}

  validator.
    On("validateUsername", "simon").
    Return(true, nil).
    On("validateUsername", "s1m0n").
    Return(false, nil).
    On("validateUsername", "anita").
    Return(true, nil)

  db := &db.MockDatabase{}

  db.
    On("Open", "dob-api.db").
    Return(nil).
    On("Close").
    Return().
    On("Get", "simon").
    Return("1980-04-20", nil).
    On("Get", "anita").
    Return("", nil)

  app := &App {
    validator: validator,
    db: db,
  }

  for _, tc := range tcs {
    fmt.Println(fmt.Sprintf("Testing username: %s",tc.Username))
    res, err := app.GetDateOfBirth(tc.Username)

    if tc.Error != nil {
      assert.Equal(t, err.Error(), tc.Error.Error())
      fmt.Println(fmt.Sprintf("PASS: Error asserted: %s",tc.Error.Error()))
    } else {
      assert.Equal(t, res, tc.Dob)
      fmt.Println(fmt.Sprintf("PASS: Result match: %s", res))
    }
  }
}

func TestUpdateUsername(t *testing.T) {

  tcs := []struct{
    Username string
    Dob string
    Error error
  }{
    {"simon", "1980-04-20", nil},
    {"s1m0n", "1980-04-20", errors.New("The username provided s1m0n didn't validate")},
    {"simon", "2099-04-20", errors.New("The date of birth provided 2099-04-20 didn't validate")},
    {"s1m0n", "2099-04-20", errors.New("The username provided s1m0n didn't validate")},
  }

  validator := &MockValidator{}

  validator.
    On("validateUsername", "simon").
    Return(true, nil).
    On("validateUsername", "s1m0n").
    Return(false, nil).
    On("validateDateOfBirth", "1980-04-20").
    Return(true, nil).
    On("validateDateOfBirth", "2099-04-20").
    Return(false, nil)

  db := &db.MockDatabase{}

  db.
    On("Open", "dob-api.db").
    Return(nil).
    On("Close").
    Return().
    On("Put", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
    Return(nil)

  app := &App {
    validator: validator,
    db: db,
  }

  for _, tc := range tcs {
    fmt.Println(fmt.Sprintf("Testing username: %s, dob: %s",tc.Username, tc.Dob))
    err := app.UpdateUsername(tc.Username, tc.Dob)

    if tc.Error != nil {
      assert.Equal(t, err.Error(), tc.Error.Error())
      fmt.Println(fmt.Sprintf("PASS: Error asserted: %s",tc.Error.Error()))
    } else {
      assert.NoError(t, err, tc.Error)
      fmt.Println("PASS: No error")
    }
  }
}
