package main

import (
	"log"
	"os"
	"strconv"

	"github.com/tkircsi/vault/models"
	"github.com/tkircsi/vault/models/mem"
	"github.com/tkircsi/vault/models/redis"
)

type application struct {
	RESTPort string
	GRPCPort string
	vault    models.VaultHandler
}

var (
	vaultdb   = "mem"
	redisAddr = ""
	redisPwd  = ""
	redisDB   = 0

	restport = ":5000"
	grpcport = ":50051"
)

func init() {
	if v, ok := os.LookupEnv("VAULT_DB"); ok {
		vaultdb = v
	}
	if v, ok := os.LookupEnv("REDIS_ADDR"); ok {
		redisAddr = v
	}
	if v, ok := os.LookupEnv("REDIS_PWD"); ok {
		redisPwd = v
	}
	if v, ok := os.LookupEnv("REDIS_DB"); ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		redisDB = i
	}

	if v, ok := os.LookupEnv("REST_PORT"); ok {
		restport = v
	}
	if v, ok := os.LookupEnv("GRPC_PORT"); ok {
		grpcport = v
	}
}

func main() {
	log.Println("vault started...")
	var v models.VaultHandler
	switch vaultdb {
	case "redis":
		v = redis.NewRedisVault(redisAddr, redisPwd, redisDB)
	default:
		v = mem.NewMemVault()
	}

	app := application{
		RESTPort: restport,
		GRPCPort: grpcport,
		vault:    v,
	}

	go app.RunREST()
	app.RunGRPC()
}
