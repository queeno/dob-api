package main

import (
  "testing"
  "os"
  "fmt"

  "github.com/stretchr/testify/assert"
)

func TestAmInLambda(t *testing.T) {
  fmt.Println("Checking if running in lambda function")
  assert.False(t, amILambda())
  fmt.Println("Now mocking running in lambda function")
  os.Setenv("AWS_EXECUTION_ENV", "SET")
  assert.True(t, amILambda())
}
