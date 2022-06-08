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

### Run command line:

```go
go run serve.go -port 8000
```


### Expected verification steps:

```go
curl localhost:8000
```

## Milestone 2
The goal of this milestone is to return the current date and time in Swatch Internet Time (
https://en.wikipedia.org/wiki/Swatch_Internet_Time ) in response to a web request.
Requirements:
- The response to a GET request to /time should be in the format in the verification steps
- There should be unit tests to verify the functionality that can be run by go test

### Install

```go
go get -u github.com/rhamitarora/assessment/milestone2
```

### Run command line:

```go
go run main.go
```

## Milestone 3
The goal of this milestone is to serve a Websockets connection that provides this information.
Requirements:
- Any WebSockets library is fine, but https://github.com/gorilla/websocket is preferred unless you
have one you prefer
- The WebSocket connection should be on /timews
- When the WebSocket connection is instantiated, it should send the current date/time as before.
- When the date/time value would change (e.g. it goes from @67.2 to @67.3), an update should be
pushed over the open WebSocket connection.
- There should be unit tests to verify the functionality that can be run by go test.

### Install

```go
go get -u github.com/rhamitarora/assessment/milestone3
```

### Run client command line:

```go
go run client/client.go
```

### Run server command line:

```go
go run server/server.go
```

## Milestone 4
The goal of this milestone is to serve a web page that has the current date/time in it, updated over
WebSockets.
Requirements:
- The server must not read from the filesystem to get the HTML/JS/CSS.
- The page should be served at /timeupdating .
- The page should work on a modern browser (Firefox, Chrome, etc)
- Anything more than text is out of scope.
### Install

```go
go get -u github.com/rhamitarora/assessment/milestone4
```
### Run server command line:

```go
go run main.go
```