package database

import "time"

type Store interface {
	Increment(key string) (int64, error)
	Expire(key string, duration time.Duration) error
	Set(key, value string, duration time.Duration) error
	Get(key string) (string, error)
	FlushDB() error
}
