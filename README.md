# Golang Skills Assessment

## Milestone 1
The goal of this milestone is to write a web server that returns a “Hello World!” message.
Requirements:
- Must use Go Modules
- Returns “Hello, world!” on a GET to the web server root (/)
- Has a command line that allows the user to provide the port that the webserver should serve on. The
server should always listen on localhost.

### Install

```go
go get -u github.com/rhamitarora/assessment/milestone1
```

### Expected command line:

```go
go run cmd/serve.go -port 8000
```


### Expected verification steps:

```go
curl localhost:8000
```


## Based On

[IPv4 Subnet Calculator ( PHP )](https://github.com/markrogoyski/ipv4-subnet-calculator-php)

## License

IPv4 Subnet Calculator (GO) is licensed under the MIT License.
