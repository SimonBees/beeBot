# beeBot Server

基于 Go 和 Fiber 框架的企业级服务端系统

## 项目结构

```
beeBot/
├── cmd/
│   └── server/
│       └── main.go          # 主入口文件
├── config/
│   └── config.go           # 配置管理
├── handlers/               # HTTP请求处理器
│   └── base_handlers.go
├── routes/                 # 路由定义
│   └── routes.go
├── models/                 # 数据模型
├── services/               # 业务逻辑层
├── utils/                  # 工具函数
├── api/                    # API定义
├── go.mod
├── go.sum
└── Makefile
```

## 功能特性

- 基于 Fiber 框架的高性能 Web 服务器
- 配置管理
- 健康检查端点
- 模块化路由结构
- 可扩展的架构

## 快速开始

### 环境要求

- Go 1.18+

### 安装依赖

```bash
go mod tidy
```

### 构建项目

```bash
make build
```

### 运行服务

```bash
make run
```

### 测试接口

- `GET /` - 欢迎页面
- `GET /health` - 健康检查
- `GET /api/v1/` - API接口组

## 环境变量

- `PORT` - 服务端口 (默认: 3000)
- `ENV` - 环境 (默认: development)
- `DB_HOST` - 数据库主机
- `DB_PORT` - 数据库端口
- `DB_USER` - 数据库用户
- `DB_PASS` - 数据库密码
- `DB_NAME` - 数据库名称
- `JWT_SECRET` - JWT密钥

## 开发规范

- 所有代码遵循 Go 语言规范
- 使用标准的 Go 项目结构
- 统一的错误处理机制
- 完整的单元测试覆盖

## 部署

```bash
# 生产环境构建
make build-linux

# 运行
./bin/server-linux
```

## 许可证

见 LICENSE 文件