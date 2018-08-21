## goRpc学习
* grpc : 基于rpc的二进制传输协议。 grpc内部有一个内部的数字模拟语言protobuf 自定义接口格式。

### 使用流程 ：
* step1.创建服务定义，创建dhcpd.proto 文件
```go
syntax = "proto3";
 // 服务命名空间
package go.micro.srv.consignment;

// 暴露给其他服务的方法
service ShippingService {  
  rpc CreateConsignment(Consignment) returns (Response) {}
}

message Consignment {
  string id = 1;
  string description = 2;
  int32 weight = 3;
  repeated Container containers = 4;
  string vessel_id = 5;
}

// 消息类型
message Container {
  string id = 1;
  string customer_id = 2;
  string origin = 3;
  string user_id = 4;
}

message Response {
  bool created = 1;
  Consignment consignment = 2;
}

```
* 消息通过protobuf处理，服务通过grpc的protobuf插件处理，把消息编译成代码，从而进行交互，protobuf 定义的结构，可以通过客户端接口，自动生成相应语言的二进制数据和功能
* step2  给服务创建一个MakeFile。(需要下载protobuf的编译器protoc), 切换到dhcpd.proto文件下， 执行 `protoc --go_out=. *.proto
` 会生成dhcpd.pb.go的新文件，这里使用了 gRPC/protobuf 库，把自定义的 protobuf 结构自动转换成你想要的代码
* step3 创建main方法
```go






```
 ### 遇到问题
 >> https://github.com/google/go-genproto