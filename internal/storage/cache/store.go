package cache

import (
	"context"
	"encoding/json"
	"go-simple/internal/config"
	"time"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	client *redis.Client
	prefix string
}

func NewCacheStore(client *redis.Client, cfg *config.Config) *Store {
	return &Store{
		client: client,
		prefix: cfg.Redis.Prefix,
	}
}

// Set stores any serializable value with a TTL, ttl is in minutes
func (r *Store) Set(ctx context.Context, key string, value interface{}, ttl int) error {
	ttlDuration := time.Duration(ttl) * time.Minute
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, data, ttlDuration).Err()
}

// Get retrieves a value and un marshals it into dest (must be a pointer)
func (r *Store) Get(ctx context.Context, key string, dest interface{}) error {
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

// Delete removes a key from the cache
func (r *Store) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// Exists checks if a key exists
func (r *Store) Exists(ctx context.Context, key string) (bool, error) {
	count, err := r.client.Exists(ctx, key).Result()
	return count > 0, err
}
