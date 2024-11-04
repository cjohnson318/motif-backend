# Motif Backend

This is a proof of concept for an HTTP server written in Go, using the Gin
framework. At the moment, this server listens for a GET request, and then uses
gRPC to call a function on a Python gRPC server, and returns a JSON holding the
Python server's response.

## Running the Code   

In one server, run `go run cmd/main.go`. This will start up an HTTP server on
port 50052. Make a GET request to `localhost:50052` and this should print
`Hello, Dork!`, or similar.

