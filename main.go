package main

import (
  "os"

  "github.com/queeno/dob-api/api"
  "github.com/queeno/dob-api/lambda"

  ll "github.com/aws/aws-lambda-go/lambda"
)

func amILambda() bool{
  if lookup, _ := os.LookupEnv("AWS_EXECUTION_ENV"); lookup != "" {
    return true
  }
  return false
}

func getDBPath() string {
  if len(os.Args) > 1 {
    return os.Args[1]
  }
  return "dob-api.db"
}

func main() {
  if amILambda() {
    ll.Start(lambda.NewLambda().HandleRouteRequest)
  } else {
    api := api.NewApi(getDBPath())
    os.Exit(api.RunServer())
  }
}
