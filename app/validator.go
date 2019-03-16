package app

import (
  "regexp"
  "time"
)

type validator interface {
  validateUsername(string) (bool, error)
  validateDateOfBirth(string) (bool, error)
}

type DobValidator struct {
  validator validator
}


func (v DobValidator) validateUsername(username string) (bool, error) {
  return regexp.MatchString(`^[a-zA-Z]+$`, username)
}

func (v DobValidator) validateDateOfBirth(dob string) (bool, error) {
  // Match format
  matched, err := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, dob)
  if !matched || err != nil {
    return false, err
  }

  //Check it is before today
  dateStamp, err := time.Parse("2006-01-02", dob)
  if err != nil {
    return false, err
  }

  today := time.Now()

  if dateStamp.After(today){
    return false, nil
  }

  return true, nil
}
