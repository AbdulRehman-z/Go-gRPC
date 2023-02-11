# Basic Implementation of gRPC in Go

## What is gRPC?
gRPC is a modern open source high performance Remote Procedure Call (RPC) framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.

## gRPC vs REST
gRPC is a modern RPC framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.

## What is this repo?
This repo is a basic implementation of gRPC in Go. It has implementations of all Unary RPCs, Server Streaming RPCs, Client Streaming RPCs and Bi-Directional Streaming RPCs.

### How to work with this repo
1. Clone this repo
2. Run `go mod tidy` to install all the dependencies
3. Run `go run server/server.go` to start the server
4. Run `go run client/client.go` to start the client
5. You can also run `go run server/server.go` and `go run client/client.go` in different terminals to see the output
6. First, you must run the server and then the client

## References
1. [gRPC](https://grpc.io/)
2. [gRPC Basics - Go](https://grpc.io/docs/languages/go/basics/)
3. [gRPC Go Quickstart](https://grpc.io/docs/languages/go/quickstart/)

## That's it!
If you have any questions, feel free to reach out to me on [Twitter](https://twitter.com/AbdulRe28323982).

