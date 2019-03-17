package db

type Database interface {
  Open() error
  Close()
  Put(string, string) error
  Get(string) string
}
