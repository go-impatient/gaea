# Gaea['dʒi:ə] 盖亚, 一个轻量级RPC业务框架

基于 [Sniper] 轻量级业务框架 (https://github.com/go-kiss/sniper) 生成的Gaea Layout.

## 系统要求

1. 类 UNIX 系统
2. go v1.12+
3. [protoc](https://github.com/google/protobuf)
4. [protoc-gen-go](https://github.com/golang/protobuf/tree/master/protoc-gen-go)

## 目录结构

```
├── cmd         # 服务子命令
├── dao         # 数据访问层
├── main.go     # 项目总入口
├── rpc         # 接口描述文件
├── server      # 控制器层
├── service     # 业务逻辑层
├── sniper.toml # 配置文件
└── util        # 业务工具库
```

## 快速入门

- [定义接口](./rpc/README.md)
- [实现接口](./server/README.md)
- [注册服务](./cmd/server/README.md)
- [启动服务](./cmd/server/README.md)
- [配置文件](pkg/conf/README.md)
- [日志系统](pkg/log/README.md)
- [指标监控](pkg/metrics/README.md)
- [链路追踪](pkg/trace/README.md)

## 使用
### 编译 sniper 工具
```bash
go build ./cmd/sniper
```

### 安装 protoc-gen-twirp
```bash
# 首次使用需要安装 protoc-gen-twirp 工具
make cmd
```

### 重命名项目总包名
```bash
./sniper rename --package moocss.com/gaea 
```
### 初次生成 rpc 代码
```bash
go run cmd/sniper/main.go rpc --server auth --service user
```

### 生成 rpc 代码
```bash
# 针对指定服务
protoc --go_out=. --twirp_out=. echo.proto
```

# 针对所有服务
find rpc -name '*.proto' -exec protoc --twirp_out=. --go_out=. {} \;

# 建议直接使用框架提供的 make 规则
make rpc
```

### 测试
```bash
curl --request "POST" \
    --header "Content-Type: application/json" \
    --data '{"msg": "Hello World"}' \
    http://localhost:8080/api/helloworld.v1.Helloworld/Echo
```

```bash
echo 'msg:"Hello World"' \
    | protoc --encode helloworld.v1.HelloworldEchoReq ./rpc/example/v1/helloworld.proto \
    | curl -s --request POST \
      --header "Content-Type: application/protobuf" \
      --data-binary @- \
      http://localhost:8080/api/helloworld.v1.Helloworld/Echo \
    | protoc --decode helloworld.v1.HelloworldEchoResp ./rpc/example/v1/helloworld.proto
```

### 批量修改文件内容
```bash
grep -rl "检索内容" --include="*" ./ | xargs sed -i "" "s/检索内容/修改后内容/g"

```

### redis
```
redis-server # 启动redis服务端
redis-cli    # 启动redis客户端
shutdown     # 关闭redis服务端
```