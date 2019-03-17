package app

import (
  "testing"
  "fmt"
  "errors"
  "time"

  "github.com/queeno/dob-api/db"

  "github.com/stretchr/testify/mock"
  "github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T){
  db := &db.MockDatabase{}
  app := NewApp(db)

  fmt.Println("Testing app has been correctly initialised")
  assert.Equal(t, app.db, db)
  fmt.Println("PASS: App correctly initialised")
}

func TestGetDateOfBirth(t *testing.T) {

  currentDate := time.Now().Format("2006-01-02")
  simonsBirthday, err := time.Parse("2006-01-02", "2099-04-20")
  if err != nil {
    t.Fatal(err)
    return
  }

  today := time.Now()
  hoursRemaining := simonsBirthday.Sub(today).Hours()
  daysRemaining := int64(hoursRemaining / 24)

  tcs := []struct{
    Username string
    Message string
    Error error
  }{
    {"simon", fmt.Sprintf("Hello, simon! Your birthday is in %d day(s)", daysRemaining), nil},
    {"anita", "Hello, anita! Happy birthday!", nil},
    {"jane", "", errors.New("")},
    {"s1m0n", "1980-04-20", errors.New("The username provided s1m0n didn't validate")},
  }

  validator := &MockValidator{}

  validator.
    On("validateUsername", "simon").
    Return(true, nil).
    On("validateUsername", "jane").
    Return(true, nil).
    On("validateUsername", "s1m0n").
    Return(false, nil).
    On("validateUsername", "anita").
    Return(true, nil)

  db := &db.MockDatabase{}

  db.
    On("Open").
    Return(nil).
    On("Close").
    Return().
    On("Get", "simon").
    Return("2099-04-20", nil).
    On("Get", "anita").
    Return(currentDate, nil).
    On("Get", "jane").
    Return("", nil)

  app := &App {
    validator: validator,
    db: db,
  }

  for _, tc := range tcs {
    fmt.Println(fmt.Sprintf("Testing username: %s",tc.Username))
    res, err := app.GetDateOfBirth(tc.Username)

    if err != nil {
      assert.Error(t, tc.Error, err)
      fmt.Println(fmt.Sprintf("PASS: Error asserted: %s",err))
    } else {
      assert.Equal(t, tc.Message, res)
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
    On("Open").
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

    if err != nil {
      assert.Error(t, tc.Error, err)
      fmt.Println(fmt.Sprintf("PASS: Error asserted: %s",err))
    } else {
      assert.NoError(t, tc.Error, err)
      fmt.Println("PASS: No error")
    }
  }
}
