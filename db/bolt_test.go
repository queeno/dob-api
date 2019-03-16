package db

import (
  "testing"
  "bytes"
  "io/ioutil"
  "path/filepath"
  "os"

  "github.com/stretchr/testify/assert"
)

func TestCreateBoltDB(t *testing.T) {
  dbFileName := t.Name() + ".dbtest"
  dbGoldenFileName := t.Name() + ".golden"

  // Create DB
  boltDB := new(BoltDB)
  boltDB.Open(dbFileName)
  defer boltDB.Close()

  gf, err := ioutil.ReadFile(filepath.Join("testdata", dbGoldenFileName))
  if err != nil {
    t.Fatal(err)
  }

  tf, err := ioutil.ReadFile(dbFileName)
  if err != nil {
    t.Fatal(err)
  }

  err = os.Remove(dbFileName)
  if err != nil {
    t.Fatal(err)
  }

  assert.True(t, bytes.Equal(tf, gf))
}

func TestStoreAndGetBoldData (t *testing.T){
  dbFileName := t.Name() + ".dbtest"

  // Create DB
  boltDB := new(BoltDB)
  boltDB.Open(dbFileName)

  testData := map[string]string{
    "simon": "1988-05-21",
    "kolja": "2001-04-20",
    "Khristine": "1004-30-12",
    "Carl": "",
  }

  for key, value := range testData {
    // Populates the DB
    err := boltDB.Store(key, value)

    if err != nil {
      t.Fatal(err)
    }
  }

  for key, value := range testData {
    result := boltDB.Get(key)
    assert.Equal(t, result, value)
  }

  err := os.Remove(dbFileName)
  if err != nil {
    t.Fatal(err)
  }
}


func TestGetBoldData (t *testing.T) {
  dbGoldenFileName := t.Name() + ".golden"

  // Create DB
  boltDB := new(BoltDB)
  boltDB.Open(filepath.Join("testdata", dbGoldenFileName))

  defer boltDB.Close()

  testData := map[string]string{
    "simon": "1988-05-21",
    "kolja": "2001-04-20",
    "Khristine": "1004-30-12",
    "Carl": "",
  }

  for key, value := range testData {
    result := boltDB.Get(key)
    assert.Equal(t, result, value)
  }

  testWrongData := map[string]string{
    "simon": "1004-10-12",
    "anita": "2012-10-11",
  }

  for key, value := range testWrongData {
    result := boltDB.Get(key)
    assert.NotEqual(t, result, value)
  }
}
