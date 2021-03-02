package main

import (
	"fmt"

	"github.com/tkircsi/vault/models"
	"github.com/tkircsi/vault/models/mem"
)

type application struct {
	port  string
	vault models.VaultHandler
}

func main() {
	fmt.Println("Service started...")
	// v := redis.NewRedisVault("", "", 0)
	v := mem.NewMemVault()
	app := application{
		port:  ":5000",
		vault: v,
	}

	app.RunREST()

	// key := "DOES_NOT_EXISTS"
	// v, err := app.vault.Get(key)
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Printf("Get value of token %q: %s\n", key, v.Value)
	// }

	// v, err = app.vault.Put(`{ "id": 444, "name": "Rambo"}`, 20*time.Minute)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ret, _ := app.vault.Get(v.Token)
	// fmt.Println(ret.Value)
}
