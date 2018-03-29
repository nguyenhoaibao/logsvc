# Log Service

A simple log service that allows to store and retrieve log events such as client ip, server ip, tags (key-value pairs) and message.

## Getting Started

### Prerequisite

  - Go >= 1.9
  - [Protocol Buffers](https://github.com/google/protobuf)

### Installing

#### Using Docker

- Start `docker-compose`

```
$ docker-compose build
$ docker-compose up -d
```

- Run DB migration

```
$ cd ./migrations
$ go run *.go [--dbaddr ...] [--dbuser ...] [--dbpasswd ...]
// Default values:
//   - dbaddr: 127.0.0.1:5432
//   - dbuser: postgres
//   - dbpasswd: mypostgrespw
```

- Run the client

```
$ cd cmd/client
$ go run main.go [--addr ...]
// Default value:
//   - addr: 127.0.0.1:8080
```

- Checkout the output.

- Cleanup

```
$ docker-compose down
```

## Development

- Install packages

```
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

- Install dependencies

```
$ go get -u github.com/golang/dep/cmd/dep
$ dep ensure
```

- Start PostgresDB service (using [docker](https://hub.docker.com/_/postgres/))

```
$ docker run --name postgres-for-logsvc -p 5432:5432 -e POSTGRES_PASSWORD=mypostgrespw -d postgres
```

- Run DB migration

```
$ cd ./migrations
$ go run *.go [--dbaddr ...] [--dbuser ...] [--dbpasswd ...]
// Default values:
//   - dbaddr: 127.0.0.1:5432
//   - dbuser: postgres
//   - dbpasswd: mypostgrespw
```

- Start the server

```
$ go run cmd/server/main.go
```

- Start the client (in new terminal)

```
$ go run cmd/client/main.go
```

- Checkout the output
- Cleanup

```
$ docker rm -f postgres-for-logsvc
```

## Running the tests

```
$ make <test|test-integration>
```
