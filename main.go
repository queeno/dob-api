package main

import (
  "github.com/queeno/dob-api/api"
  "os"
)

func main() {
  api := api.NewApi("dob-api.db")
  os.Exit(api.RunServer())
}
