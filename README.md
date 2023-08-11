
# pubsub_poc

## Prerequisites

* Go 1.20
* Docker
* Running RabbitMQ instance

```
$ docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

## Server
#### Build
```
$ go build
```
#### Run
```
$ ./pubsub_poc
```

## CLI client
#### Build
```
$ go build ./cmd/cli_client
```
#### Run
```
$ ./cmd/cli_client/cli_client -h
```
#### Usage: add item
```
$ ./cmd/cli_client/cli_client add --key=key1 --val=val1
```
#### Usage: remove item
```
$ ./cmd/cli_client/cli_client remove --key=key1
```
#### Usage: get item
```
$ ./cmd/cli_client/cli_client get --key=key1
```
#### Usage: get all items
```
$ ./cmd/cli_client/cli_client getall
```

## Testing
```
$ go test ./...
```

