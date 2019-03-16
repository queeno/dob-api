package db

type Database interface {
  Open(string) error
  Close()
  Put(string, string) error
  Get(string) (string, error)
}
