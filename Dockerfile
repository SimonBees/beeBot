# 使用官方Golang镜像作为构建环境
FROM golang:1.22-alpine AS builder

# 安装git（构建依赖可能需要）
RUN apk add --no-cache git

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server cmd/server/main.go

# 使用轻量级alpine镜像作为运行环境
FROM alpine:latest

# 安装ca-certificates包，用于HTTPS请求
RUN apk --no-cache add ca-certificates

# 创建非root用户
RUN adduser -D -s /bin/sh beeuser

# 设置工作目录
WORKDIR /app

# 从builder阶段复制编译好的二进制文件
COPY --from=builder /app/server .

# 更改文件所有者
RUN chown beeuser:beeuser server

# 切换到非root用户
USER beeuser

# 暴露端口
EXPOSE 3000

# 启动命令
CMD ["./server"]