package main

import (
  "os"

  "github.com/queeno/dob-api/api"
  thisLambda "github.com/queeno/dob-api/lambda"

  "github.com/aws/aws-lambda-go/lambda"
)

func mainLambda(){
  myLambda := &thisLambda.Lambda{}
  lambda.Start(myLambda.HandleRouteRequest)
}

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
