package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")
)

type Token struct {
	Token string
	Value string
}

type VaultHandler interface {
	Get(string) (*Token, error)
	Put(string, time.Duration) (*Token, error)
}
