package redis

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/tkircsi/vault/models"
	"github.com/tkircsi/vault/services"
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
	v, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	if v != "PONG" {
		log.Fatal("can not get PONG from redis")
	}
	log.Printf("connected to redis OK")
	return &RedisVault{rdb: rdb}
}

func (r *RedisVault) Get(key string) (*models.Token, error) {
	cipher, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, models.ErrNoRecord
	}
	val, err := services.Decrypt(cipher)
	if err != nil {
		return nil, err
	}
	return &models.Token{Token: key, Value: val}, nil
}

func (r *RedisVault) Put(val string, exp time.Duration) (*models.Token, error) {
	cipher, err := services.Encrpyt(val)
	if err != nil {
		return nil, err
	}
	_, err = r.rdb.Set(ctx, cipher, cipher, exp).Result()
	if err != nil {
		return nil, err
	}
	return &models.Token{Token: cipher, Value: cipher}, nil
}
