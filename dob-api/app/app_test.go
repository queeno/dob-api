package app

import (
  "testing"
  "fmt"
  "errors"
  "time"

  "github.com/queeno/dob-api/dob-api/db"

  "github.com/stretchr/testify/mock"
  "github.com/stretchr/testify/assert"
)

func TestDaysRemainingToNextBirthday(t *testing.T){
  tcs := []struct{
    DateOfBirth time.Time
    Today time.Time
    DaysRemaining int
  }{
    {time.Date(1975, 12, 21, 0, 0, 0, 0, time.UTC), time.Date(2011, 12, 31, 0, 0, 0, 0, time.UTC), 356},
    {time.Date(1988, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2011, 12, 21, 0, 0, 0, 0, time.UTC), 10},
    {time.Date(7, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2010, 12, 21, 0, 0, 0, 0, time.UTC), 10},
    {time.Date(2004, 2, 29, 0, 0, 0, 0, time.UTC), time.Date(2012, 12, 21, 0, 0, 0, 0, time.UTC), 70},
    {time.Date(2004, 2, 29, 0, 0, 0, 0, time.UTC), time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC), 59},
    {time.Date(2004, 2, 29, 0, 0, 0, 0, time.UTC), time.Date(2012, 2, 29, 0, 0, 0, 0, time.UTC), 0},
  }

  for _, tc := range tcs {
    app := &App {
      today: tc.Today,
    }

    fmt.Println(fmt.Sprintf("Today is: %s. DateOfBirth is: %s. No of days: %d", tc.Today, tc.DateOfBirth, tc.DaysRemaining))
    drm := app.daysRemainingToNextBirthday(tc.DateOfBirth)
    assert.Equal(t, tc.DaysRemaining, drm)
    fmt.Println("PASS: Matched")
  }
}

func TestNewApp(t *testing.T){
  db := &db.MockDatabase{}
  app := NewApp(db)

  fmt.Println("Testing app has been correctly initialised")
  assert.Equal(t, app.db, db)
  fmt.Println("PASS: App correctly initialised")
}

func TestGetDateOfBirth(t *testing.T) {

  tcs := []struct{
    Username string
    Today time.Time
    Birthday string
    Message string
    Error error
  }{
    {"simon", time.Date(2099, 4, 10, 0, 0, 0, 0, time.UTC), "2099-04-20", "Hello, simon! Your birthday is in 10 day(s)", nil},
    {"anita", time.Date(2099, 4, 20, 0, 0, 0, 0, time.UTC), "2099-04-20", "Hello, anita! Happy birthday!", nil},
    {"jane", time.Date(2099, 4, 10, 0, 0, 0, 0, time.UTC), "2099-04-20", "", errors.New("")},
    {"s1m0n", time.Date(1988, 4, 10, 0, 0, 0, 0, time.UTC), "2099-04-10", "", errors.New("")},
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
    Return("2099-04-20",nil).
    On("Get", "anita").
    Return("2099-04-20",nil).
    On("Get", "jane").
    Return("",nil).
    On("Get", "s1m0n").
    Return("2099-04-10",nil)

  for _, tc := range tcs {
    app := &App {
      validator: validator,
      db: db,
      today: tc.Today,
    }

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
    today: time.Now(),
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
