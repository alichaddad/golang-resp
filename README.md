# Golang RESP server

A simple TCP implementation of the Redis [RESP protocol](https://redis.io/docs/reference/protocol-spec/).
This implementation is done for learning purposes and it supports only the following two custom commands:

- `myping`: pings the server
- `testurl`: checks if the given url is reachable

## Using the commands

In a redis shell run you can try out the commands as follows:

For `myping`.

```bash
127.0.0.1:7000> myping
"pong"
```

For `testurl`.

```bash
127.0.0.1:7000> testurl https://google.com
"true"
```

## Running the Server

```bash
go run main.go --address localhost:7000
```
