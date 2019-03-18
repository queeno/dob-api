package main

import (
  "os"

  "github.com/queeno/dob-api/dob-api/api"
  thisLambda "github.com/queeno/dob-api/dob-api/lambda"

  "github.com/aws/aws-lambda-go/lambda"
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
    myLambda := &thisLambda.Lambda{}
    lambda.Start(myLambda.HandleRouteRequest)
  } else {
    api := api.NewApi(getDBPath())
    os.Exit(api.RunServer())
  }
}
