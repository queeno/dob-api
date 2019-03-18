package db

import (
  "time"

	"github.com/boltdb/bolt"
)

type BoltDB struct {
  db *bolt.DB
  FilePath string
}

func (blt BoltDB) Get(key string) (string, error) {
  value := make([]byte, 0)
  err := blt.db.View(func(tx *bolt.Tx) error {
    bucket := tx.Bucket([]byte("DateOfBirths"))
    if bucket == nil {
      return nil
    }
	  value = bucket.Get([]byte(key))
    return nil
  })

  return string(value), err
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

func (blt *BoltDB) Open() error {
	db, err := bolt.Open(blt.FilePath, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return err
	}

	blt.db = db

	return nil
}
