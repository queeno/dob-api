package app

import (
  "errors"
  "fmt"

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
}

func (a *App) Initialise(db db.Database) {
  a.db = db
  a.validator = &dobValidator{}
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

  err = a.db.Open("dob-api.db")
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

  err = a.db.Open("dob-api.db")
  if err != nil {
    return "", err
  }
  defer a.db.Close()

  dob, err := a.db.Get(username)
  if err != nil {
    return "", err
  }

  return dob, nil
}
