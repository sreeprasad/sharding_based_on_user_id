### Sharding based on user id

There is only 2 databases.
If user ID%2 == 0 then choose database 1 other wise use database 2

## How to run

start the docker to run the 2 postgres databases

```shell
docker-compose up --build
```

start the server in another terminal

```shell
go run simulate_sharing_in_server.go
```

start curl for different user ids
If user id is even, then request should go to database 1
Otherwise request should go to database 2.

Curl with even user id

```shell
 curl   "http://localhost:8080/execute?userid=6"
```

Curl with odd user id

```shell
 curl   "http://localhost:8080/execute?userid=5"
```
