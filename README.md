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

## 三种心跳检测定时器

1. 检测超过最大空闲时间
2. 检测超过最大连接时间
3. 指定时间段内没有有给数据交互
   以上三种任意三种满足都应该断开连接

| 类型          | 检测粒度  | 用途         | 建议默认值      |
|-------------|-------|------------|------------|
| 最大空闲时间      | 分钟级   | 回收长时间无操作连接 | 1-5 分钟     |
| 最大连接时间      | 小时-天级 | 负载均衡、强制重连  | 12-48 小时   |
| 心跳超时（短期无交互） | 秒级    | 快速检测异常     | 心跳间隔 + 5 秒 |

## 实现的步骤

按照下面tag版本可以一步步观察实现步骤，tag之间对比观察更明显
建议将tag拉取到本地，用IDE工具进行版本之间对比参照

## kafka

```shell
kafka-console-producer.sh --broker-list 127.0.0.1:9092 --topic msgChatTransfer
```

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
    * v7.0.4 mysql数据读写操作 ✅
    * v7.0.5 api中间件的使用 ✅
8. v8.0.0 实现user rpc/web服务 ✅
    * v8.0.1 增加rpc用户手机号、密码加密、雪花id、jwt、注册登录的业务 ✅
    * v8.0.2 完善user api服务、完成登录注册功能、用户详情 ✅
    * v8.0.3 优化响应输出，api错误码的统一处理 ✅
9. v9.0.0 实现im服务用户登入连接，鉴权(将im/ws/websocket/authentication.go按jwt方式替换)
10. v10.0.0 实现im心跳检测
    * v10.0.1 区分心跳消息和普通消息，并优化
    * v10.0.2 使用带心跳检测的连接
11. v11.0.0 好友私聊，私聊数据存储、请求信息、实现私聊
12. v12.0.0 使用kafka构建异步消费服务
13. v13.0.0 基于kafka异步数据存储落地及消息通信
    * v13.0.1 构建好websocket客户端
    * v13.0.2 超级token验证，mq中的服务业务
    * v13.0.3 push消息到客户端，websocket将接收消息写入消息队列
14. v14.0.0 实现消息的ack机制,基础结构
    * v14.0.1 ack options配置与消息属性
    * v14.0.2 实现ack机制
15. v15.0.0 用户拉取离线消息
    * v15.0.1 完成im-rpc服务的功能开发
    * v15.0.2 任务新增消息记录
    * v15.0.3 完成im-api服务的功能开发




