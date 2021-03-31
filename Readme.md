# gRPC/REST Vault v0.0.2

Vault is a simple microservice, which is exchanging sensitive data with a token. It is a good use-case, when we do not want to share sensitive data with the client, but data is needed for other back-end services. Data is stored with AEAD(authenticated encryption with associated data) with AES-GCM.
The storage can be in-memory or redis. With Redis backend we can specify an expiration time for the token and data. The service is available via REST API/JSON or gRPC protocol.

## Environment variable

#### VAULT_DB

Can be 'redis' or 'mem'.

#### REDIS_ADDR

The address of the Redis service in format 'host:port'

#### REDIS_PWD

The default Redis user password

#### REDIS_DB

An integer, the default is 0 (default db)

#### REST_PORT

The REST/HTTP port. Default is ":5000"

#### GRPC_PORT

gRpc port. The default is ":500051"

#### GIN_MODE

The REST/HTTP server mode. It can be 'debug' or 'release'. Use 'release' for production.

#### SECRET_KEY

The 32 bytes length AES-256 secret key

### Example docker-compose file

```
version: '2.4'

services:
  vaultsrv:
    build: .
    ports:
      - "5000:5000"
      - "50051:50051"
    environment:
      - VAULT_DB=redis
      - REDIS_ADDR=vaultdb:6379
      - REDIS_PWD=mypassword
      - REDIS_DB=0
      - REST_PORT=:5000
      - GRPC_PORT=:50051
      - GIN_MODE=debug # 'release' for production
      - SECRET_KEY=32-bytes-length-secret-key-123-4
    depends_on:
      - vaultdb

  vaultdb:
    image: redis:6.2
    command: redis-server /usr/local/etc/redis/redis.conf
    ports:
      - "6379:6379"
    volumes:
      - $PWD/redis.conf:/usr/local/etc/redis/redis.conf

```

## Services

### REST

Get service requires a token as input parameter and returns a token and value object.

#### Get

Request

```
http://localhost:5000/v1/get/ad0GK_NGTOUInNOcJWghvHu7SRjOYe-RNi2h_XHJTU3TS0Vm2xEcjQH8LRcIOPRoNoqgwuWLm-5NDQ==
```

Response

```
{
  "token": "ad0GK_NGTOUInNOcJWghvHu7SRjOYe-RNi2h_XHJTU3TS0Vm2xEcjQH8LRcIOPRoNoqgwuWLm-5NDQ==",
  "value": "sensitive data"
}
```

#### Put

Put service stores a sensitive value for a duration and returns a token. It requres a string(value) and an integer (seconds) as input parameters and returns a string token as result or an error.

Request

```
http://localhost:5000/v1/put
Content-Type=application/json
```

Body

```
{
    "value": "{ \"id\": \"ABC_777888\" }",
    "expire": 30
}
```

Response

```
{
    "token": "ad0GK_NGTOUInNOcJWghvHu7SRjOYe-RNi2h_XHJTU3TS0Vm2xEcjQH8LRcIOPRoNoqgwuWLm-5NDQ=="
}
```

### gRPC (Get/Put)

```
func testGRPC() {
	log.Println("GRPC client started....")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	cl := ClientData{
		Name:    "Kiss Elemér",
		ID:      1,
		Email:   "elem@mail.hu",
		Account: "111-1111-2222",
	}
	b, _ := json.Marshal(cl)

	c := vaultpb.NewVaultServiceClient(cc)
	req := &vaultpb.PutRequest{
		Value:  string(b),
		Expire: 60,
	}
	res, err := c.Put(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PutResponse: %s\n", res.Token)

	gr := &vaultpb.GetRequest{
		Token: res.Token,
	}
	gres, err := c.Get(context.Background(), gr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GetResponse: %s\n", gres.Value)
}
```
