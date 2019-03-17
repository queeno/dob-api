package main

import (
  "github.com/queeno/dob-api/api"
  "os"
)

func getDBPath() string {
  if len(os.Args) > 1 {
    return os.Args[1]
  }
  return "dob-api.db"
}

func main() {
  api := api.NewApi(getDBPath())
  os.Exit(api.RunServer())
}
