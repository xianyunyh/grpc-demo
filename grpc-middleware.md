grpc-middlerware 是gRPC的interceptor，可以对RPC的请求和响应进行拦截处理，而且既可以在客户端进行拦截，也可以对服务器端进行拦截。
主动有四种拦截器
- 服务端单向拦截器UnaryServerInterceptor 
    ```go 
    func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
    ```
- 服务的流拦截器 StreamServerInterceptor
    ```go
    type StreamServerInterceptor func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error
    ```
- 客户端单向拦截器 UnaryClientInterceptor 
    ```go
    type UnaryClientInterceptor func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error

    ```
- 客户端流拦截器 StreamClientInterceptor 
    ```go
    type StreamClientInterceptor func(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, streamer Streamer, opts ...CallOption) (ClientStream, error)
    ```

社区有很多实现
- https://github.com/grpc-ecosystem/go-grpc-middleware

