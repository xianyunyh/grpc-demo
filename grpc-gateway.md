
grpc-gatway是一个插件，把http RESTful API的请求转到gRPC处理。
## 安装

```shell
$ go get github.com/grpc-ecosystem/grpc-gateway/v2
$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```


## 定义接口

```proto
import "google/api/annotations.proto";
service Post {
    rpc ListPost(ListPostReuquest) returns(ListPostResponse){
        option (google.api.http) = {
			    get: "/posts"
		    }; 
    };
    rpc PostDetail(GetPostOne) returns (PostItem) {
        option (google.api.http) = {
            get: "/posts/{id}"
        };
    };
}

```
## 生成代码

```shell
$ protoc -I=. --go_out=./models \
  --go-grpc_out=./models \
  --grpc-gateway_out=./models \
   --grpc-gateway_opt=logtostderr=true  \
   --grpc-gateway_opt=generate_unbound_methods=true  \
    --openapiv2_out=./openapiv2 \
    --openapiv2_opt=logtostderr=true \
   proto/hello.proto

```