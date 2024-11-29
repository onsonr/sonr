package store

type Store interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}
