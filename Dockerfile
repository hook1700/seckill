FROM golang:1.21-alpine AS builder

# 国内代理（腾讯云必加）
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=off

WORKDIR /app

# 1️⃣ 只拷 go.mod
COPY go.mod .

# 2️⃣ 在容器里生成 go.sum
RUN go mod tidy

# 3️⃣ 下载依赖（此时 go.sum 已存在）
RUN go mod download

# 4️⃣ 拷源码
COPY . .

# 5️⃣ 编译
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app

# ========== 运行镜像 ==========
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]