package main

import (
  "github.com/queeno/dob-api/api"
  "os"
)

func main() {
  api := api.NewApi()
  os.Exit(api.RunServer())
}
