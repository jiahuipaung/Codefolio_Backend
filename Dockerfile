# 构建阶段
FROM golang:1.23.3-alpine AS builder

# 设置工作目录
WORKDIR /app

# 设置 Go 代理
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on

# 安装必要的构建工具
RUN apk add --no-cache git

# 从远程仓库克隆代码
RUN git clone https://github.com/jiahuipaung/Codefolio_Backend.git .

# 下载依赖
RUN go mod download

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/user-service -mod=mod ./internal/user/main.go

# 运行阶段
FROM alpine:latest

# 安装必要的运行时依赖
RUN apk add --no-cache ca-certificates tzdata

# 设置时区
ENV TZ=Asia/Shanghai

# 创建非 root 用户
RUN adduser -D -g '' appuser

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/user-service .
# 复制配置文件
COPY --from=builder /app/internal/common/config/global.yaml ./config/
# 复制其他必要的文件
COPY --from=builder /app/scripts ./scripts

# 设置权限
RUN chown -R appuser:appuser /app

# 切换到非 root 用户
USER appuser

# 暴露端口
EXPOSE 8080

# 启动命令
CMD ["./user-service"] 