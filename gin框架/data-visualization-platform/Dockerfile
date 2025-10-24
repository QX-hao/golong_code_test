# 构建阶段
FROM golang:1.22-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN go build -o data-visualization-platform

# 运行阶段
FROM alpine:latest

# 安装必要的运行时依赖
RUN apk --no-cache add ca-certificates

# 创建应用用户
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/data-visualization-platform .
COPY --from=builder /app/config ./config
COPY --from=builder /app/templates ./templates

# 创建必要的目录
RUN mkdir -p uploads static logs

# 设置文件权限
RUN chown -R appuser:appgroup /app

# 切换到非root用户
USER appuser

# 暴露端口
EXPOSE 9033

# 启动应用
CMD ["./data-visualization-platform"]