## 安装protoc
[下载地址](https://github.com/protocolbuffers/protobuf/releases/tag/v3.19.4)

## 安装gRPC

```shell
$ go get google.golang.org/protobuf
$ go get google.golang.org/grpc
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

## 生成代码
```
$ protoc -I=. \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```

这里简单介绍一下 proto 的编译命令:

- –proto_path或者-I ：指定 import 路径，可以指定多个参数，编译时按顺序查找，不指定时默认查找当前目录。
.proto 文件中也可以引入其他 .proto 文件，这里主要用于指定被引入文件的位置。
- -–go_out：golang编译支持，指定输出文件路径
其他语言则替换即可，比如 --java_out 等等
- –go_opt：指定参数，比如--go_opt=paths=source_relative就是表明生成文件输出使用相对路径。

- --go-grpc_out grpc支持 指定grpc输出文件路径

- path/to/file.proto ：被编译的 .proto 文件放在最后面