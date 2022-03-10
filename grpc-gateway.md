## 安装

```shell
$ go get github.com/grpc-ecosystem/grpc-gateway/v2
$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## 生成代码

```shell
$ protoc -I=. --go_out=./models \
  --go-grpc_out=./models \
  --grpc-gateway_out=./models \
   --grpc-gateway_opt=logtostderr=true  \
   --grpc-gateway_opt=generate_unbound_methods=true  \
    --openapiv2_out=./gen/openapiv2 \
    --openapiv2_opt=logtostderr=true \
   proto/hello.proto

```