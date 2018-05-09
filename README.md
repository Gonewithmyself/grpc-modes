# grpc-modes

Golang 社区最受欢迎的微服务框架 [go-micro](https://github.com/micro/go-micro) 整合了 gRPC 的四种数据交互模式，特此简要介绍，具体参考：

gRPC [原文档](https://grpc.io/docs/tutorials/basic/go.html) 与 [中文文档](https://doc.oschina.net/grpc)



## 目录结构

```
$GOPATH
└── grpc
    ├── simple			// 简单模式 RPC
    │   ├── client	
    │   │   └── client.go		# 客户端代码
    │   ├── proto			
    │   │   ├── user.pb.go	
    │   │   └── user.proto		# 通信的 protobuf 协议
    │   └── server
    │       └── server.go		# 服务端代码
    ├── server-side-streaming	// 服务端流式 RPC 
    ├── client-side-streaming	// 客户端流式 RPC 
    └── bidirectional-streaming	// 客户端与服务端双向流式 RPC
```



## UserService 微服务

本项目中定义了一个微服务：`UserService`，它只有一个 RPC：`GetUserInfo()`

```protobuf
syntax = "proto3";
package grpc.simple;

// 定义 UserService 微服务
service UserService {
    // 微服务中获取用户信息的 RPC 函数
    rpc GetUserInfo (UserRequest) returns (UserResponse);
}

// 客户端请求的格式
message UserRequest {
    int32 ID = 1;
}

// 服务端响应的格式
message UserResponse {
    string name = 1;
    int32 age = 2;
}
```

在 `GetUserInfo()` 函数中模拟了一个数据库，存储用户的姓名和年龄：

```go
// ID 为 key，用户信息为 value 模拟数据库查询结果
var users = map[int32]pb.UserResponse{
	1: {Name: "Dennis MacAlistair Ritchie", Age: 70},
	2: {Name: "Ken Thompson", Age: 75},
	3: {Name: "Rob Pike", Age: 62},
}
```

客户端请求带上 ID，查询后将用户信息作为响应返回。





## 客户端与服务端进行数据交互的四种模式

### simpe 简单模式 RPC

客户端发起一个请求到服务端，服务端返回一个响应。

client 请求 ID 为 2 的用户数据，server 返回 ID 为 2 的用户数据：

![simple](http://p7f8yck57.bkt.clouddn.com/2018-05-09-134349.gif)



### server-side streaming 服务端流式 RPC 

客户端发起一个请求到服务端，服务端返回一段连续的数据流响应。

client 请求 1 的用户数据，server 返回 1、2、3 的用户数据流：

![server-side-streaming](http://p7f8yck57.bkt.clouddn.com/2018-05-09-134746.gif)





### client-side streaming 客户端流式 RPC 

客户端将一段连续的数据流发送到服务端，服务端返回一个响应。

client 请求 1、2、3 的用户数据流，server 返回 3 的用户数据：

![client-side-streaming](http://p7f8yck57.bkt.clouddn.com/2018-05-09-135043.gif)



### Bidirectional streaming 双向数据流模式的 gRPC

客户端将连续的数据流发送到服务端，服务端返回交互的数据流。

client 依次请求 1、2、3 的用户数据流，服务端依次返回 1、2、3 的用户数据流：

![bidirectional-streaming](http://p7f8yck57.bkt.clouddn.com/2018-05-09-135326.gif)





## 最后

最近在系统的学习 Golang 的微服务，从 gRPC 开始，到 go-micro、Docker 化微服务等，每周更新。欢迎关注我的博客 [wuYinBlog](https://github.com/wuYin/blog)

希望本项目对你有所帮助 ☺️