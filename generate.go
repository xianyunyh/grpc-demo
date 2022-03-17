package main

//go:generate protoc -I=. --openapiv2_opt=logtostderr=true --openapiv2_out=.  --go_out=./models  --go-grpc_out=./models --grpc-gateway_out=./models --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=generate_unbound_methods=true   proto/post.proto
