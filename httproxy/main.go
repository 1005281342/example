package main

//go:generate protoc --go_out=plugins=grpc:. --grpchan_out=. ./hello/hello.proto
