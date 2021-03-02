package redis

import (
	"errors"
	"testing"
	"time"
)

var (
	// redis comnfiguration
	// localhost:6379 by default
	addr = ""
	pwd  = ""
	db   = 1

	data = `{ "id": 12 }`
	exp  = 60 * time.Second

	ErrTokenMismatch = errors.New("redis: token mismatch")
)

func TestNewRedisVault(t *testing.T) {
	rv := NewRedisVault(addr, pwd, db)
	r, err := rv.rdb.Ping(ctx).Result()
	if err != nil {
		t.Fatal(err)
	}
	if r != "PONG" {
		t.Fatal("redis: missing PONG")
	}
}

func TestPut(t *testing.T) {
	rv := NewRedisVault(addr, pwd, db)
	token, err := rv.Put(data, exp)
	if err != nil {
		t.Fatal(err)
	}
	if token.Value != data {
		t.Fatal(ErrTokenMismatch)
	}
}

func TestGet(t *testing.T) {
	rv := NewRedisVault(addr, pwd, db)
	token, _ := rv.Put(data, exp)

	gt, err := rv.Get(token.Token)
	if err != nil {
		t.Fatal(err)
	}
	if (gt.Token != token.Token) || (gt.Value != token.Value) {
		t.Fatal(ErrTokenMismatch)
	}
}
