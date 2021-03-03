package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/tkircsi/vault/models"
	"github.com/tkircsi/vault/models/mem"
	"github.com/tkircsi/vault/models/redis"
)

type application struct {
	port  string
	vault models.VaultHandler
}

var (
	vaultdb   = "mem"
	redisAddr = ""
	redisUsr  = ""
	redisDB   = 0
)

func init() {
	if v, ok := os.LookupEnv("VAULT_DB"); ok {
		vaultdb = v
	}
	if v, ok := os.LookupEnv("REDIS_ADDR"); ok {
		redisAddr = v
	}
	if v, ok := os.LookupEnv("REDIS_USER"); ok {
		redisUsr = v
	}
	if v, ok := os.LookupEnv("REDIS_DB"); ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		redisDB = i
	}
}

func main() {
	fmt.Println("vault service started...")
	var v models.VaultHandler
	switch vaultdb {
	case "redis":
		v = redis.NewRedisVault(redisAddr, redisUsr, redisDB)
	default:
		v = mem.NewMemVault()
	}

	app := application{
		port:  ":5000",
		vault: v,
	}

	app.RunREST()
}
