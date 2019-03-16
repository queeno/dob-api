package db

type Database interface {
  Initialise() error
  Store(string, string) error
  Get(string) (string, error)
}
