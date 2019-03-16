package db

import (
	//"log"
  "time"
	"fmt"

	"github.com/boltdb/bolt"
)

type BoltDB struct {
  db *bolt.DB
  filePath string
}

func (blt BoltDB) Get(key string) string {
  value := make([]byte, 0)
  blt.db.View(func(tx *bolt.Tx) error {
    bucket := tx.Bucket([]byte("DateOfBirths"))
	  value = bucket.Get([]byte(key))
    return nil
  })

  return string(value)
}


func (blt BoltDB) Put(key string, value string) error {
  err := blt.db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists([]byte("DateOfBirths"))
    if err != nil {
      return err
    }

    err = b.Put([]byte(key), []byte(value))
  	if err != nil {
  	  return err
    }
    return nil
  })
  return err
}

func (blt *BoltDB) Close() {
  blt.db.Close()
}

func (blt *BoltDB) Open(location string) {
	fmt.Println(location)
	db, err := bolt.Open(location, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		panic(err)
	}

	blt.db = db
	blt.filePath = location
}
