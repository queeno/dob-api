package main

import(
  "testing"
  "os"
  "syscall"

  "github.com/queeno/dob-api/dob-api/api"
)

func TestMain (m *testing.M) {
  setup()
  retCode := m.Run()
  teardown()
  os.Exit(retCode)
}

func setup() {
  api := api.NewApi("/tmp/dobAPITest.db")
  go func() {
    api.RunServer()
  }()
}

func teardown() {
  err := os.Remove("/tmp/dobAPITest.db")
  if err != nil {
    panic(err)
  }

  syscall.Kill(syscall.Getpid(), syscall.SIGINT)
  if err != nil {
    panic(err)
  }
}
