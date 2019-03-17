package app

import (
  "errors"
  "fmt"
  "time"

  "github.com/queeno/dob-api/db"
)

type MyApp interface {
  UpdateUsername(string, string) (error)
  GetDateOfBirth(string) (string, error)
  Initalise(db.Database)
}

type App struct {
  app MyApp
  validator Validator
  db db.Database
  today time.Time
}

func (a App) UpdateUsername(username string, dateofbirth string) error {
  u, err := a.validator.validateUsername(username)
  if err != nil {
    return err
  }
  if !u {
    return errors.New(fmt.Sprintf("The username provided %s didn't validate", username))
  }

  d, err := a.validator.validateDateOfBirth(dateofbirth)
  if err != nil {
    return err
  }
  if !d {
    return errors.New(fmt.Sprintf("The date of birth provided %s didn't validate", dateofbirth))
  }

  err = a.db.Open()
  if err != nil {
    return err
  }
  defer a.db.Close()

  err = a.db.Put(username, dateofbirth)
  if err != nil {
    return err
  }

  return nil
}

func (a App) GetDateOfBirth(username string) (string, error) {
  u, err := a.validator.validateUsername(username)
  if err != nil {
    return "", err
  }
  if !u {
    return "", errors.New(fmt.Sprintf("The username provided %s didn't validate", username))
  }

  err = a.db.Open()
  if err != nil {
    return "", err
  }
  defer a.db.Close()

  dob, err := a.db.Get(username)
  if err != nil {
    return "", err
  }

  dateStamp, err := time.Parse("2006-01-02", dob)
  if err != nil {
    return "", err
  }

  daysRemaining := a.daysRemainingToNextBirthday(dateStamp)

  if daysRemaining == 0 {
    return fmt.Sprintf("Hello, %s! Happy birthday!", username), nil
  }

  return fmt.Sprintf("Hello, %s! Your birthday is in %d day(s)", username, daysRemaining), nil
}

func (a App) daysRemainingToNextBirthday(dateOfBirth time.Time) int {
  displacement := 0
  nextBirthday := time.Date(a.today.Year(), dateOfBirth.Month(), dateOfBirth.Day(), 0, 0, 0, 0, time.UTC)

  if nextBirthday.Before(a.today) {
    nextBirthday = nextBirthday.AddDate(1,0,0)
    displacement = time.Date(a.today.Year(), 12, 31, 0, 0, 0, 0, time.UTC).YearDay()
  }

  return (nextBirthday.YearDay() - a.today.YearDay()) + displacement
}

func NewApp(db db.Database) *App {
  app := &App {
    validator: &dobValidator{},
    db: db,
    today: time.Now(),
  }
  return app
}
