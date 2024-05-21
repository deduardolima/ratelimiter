package database

import (
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type RedisStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStore(client *redis.Client) *RedisStore {
	return &RedisStore{client: client, ctx: context.Background()}
}

func (r *RedisStore) Increment(key string) (int64, error) {
	return r.client.Incr(r.ctx, key).Result()
}

func (r *RedisStore) Expire(key string, duration time.Duration) error {
	return r.client.Expire(r.ctx, key, duration).Err()
}

func (r *RedisStore) Set(key, value string, duration time.Duration) error {
	return r.client.Set(r.ctx, key, value, duration).Err()
}

func (r *RedisStore) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

func (r *RedisStore) FlushDB() error {
	return r.client.FlushDB(r.ctx).Err()
}
