package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rs/xid"
	"github.com/tkircsi/vault/models"
)

type RedisVault struct {
	rdb *redis.Client
}

var ctx = context.Background()

func NewRedisVault(addr string, pwd string, db int) *RedisVault {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
	return &RedisVault{rdb: rdb}
}

func (r *RedisVault) Get(key string) (*models.Token, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, models.ErrNoRecord
	}
	return &models.Token{Token: key, Value: val}, nil
}

func (r *RedisVault) Put(val string, exp time.Duration) (*models.Token, error) {
	key := xid.New().String()
	_, err := r.rdb.Set(ctx, key, val, exp).Result()
	if err != nil {
		return nil, err
	}
	return &models.Token{Token: key, Value: val}, nil
}
