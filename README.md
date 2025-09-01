# my_chat

使用golang实现的websocket服务端

## go_zero

[go-zero官网](https://go-zero.dev)

### 安装

```shell
go install github.com/zeromicro/go-zero/tools/goctl@latest
goctl --version
```

### 快速构建api/rpc服务

```shell
cd demo && mkdir userdemo && cd userdemo

# 一般不这样使用， 而是应该先创建好rpc.proto或api.api
goctl rpc new rpc
goctl api new api
```

* 构建proto文件`demo/userdemo/rpc/user.proto`
* 构建api文件`demo/userdemo/api/user.api`
```shell
goctl rpc protoc user.proto  --go_out=.  --go-grpc_out=.  --zrpc_out=.
goctl api go -api user.api -dir . -style gozero
```

mysql数据库模型的构建
```shell
goctl model mysql ddl --src user.sql --dir "./models/" -c
```
```shell
CREATE DATABASE `user` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
```

## 实现的步骤

按照下面tag版本可以一步步观察实现步骤，tag之间对比观察更明显
建议将tag拉取到本地，用IDE工具进行版本之间对比参照

1. v1.0.0 完成im基本服务框架结构搭建 ✅
2. v2.0.0 完成消息的路由分发 ✅
3. v3.0.0 存储连接对象及设计鉴权 ✅
4. v4.0.0 消息的发送 ✅
5. v5.0.0 使用options代码风格的优化、连接的鉴权、记录连接通道 ✅
6. v6.0.0 接入消息发送,路由加载 ✅
7. v7.0.0 测试go-zero使用 ✅
   * v7.0.1 rpc服务能启动，并访问 ✅
   * v7.0.2 api服务实现并能启动访问 ✅
   * v7.0.3 api服务对rpc服务的调用 ✅
   * v7.0.4 mysql数据读写操作
   * v7.0.5 api中间件的使用
8. v8.0.0 实现user rpc/web服务 
   * v8.0.1 增加rpc用户手机号、密码加密、雪花id、jwt、注册登录的业务
   * v8.0.1 完善user api服务、完成登录注册功能、用户详情
   * 优化响应输出
9. v9.0.0 实现im服务用户登入连接，鉴权

