## proto文件格式

`syntax = "proto3";` 第一行指定版本 

## 定义包名

```proto
option go_package="xxx";
```
## 定义消息

```proto
//枚举
enum SEX {
    WOMAN = 0;
    MAN = 1;;
}
message User {
    int32 id = 1;
    string name =2;//字符串
    double money  = 3; //浮点 float、double =>float32 float64
    SEX sex = 4; //枚举
    repeated string friends = 5; //定义切片 []string
    map<string,string> other = 6;//定义map map[string]string
    ok bool = 7;// 布尔
};

message UserResponse {

}
```
## 定义服务

```proto
service UserService {

    rpc UserList(User) returns (UserResponse);
}
```