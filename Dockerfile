FROM golang:1.21-alpine AS builder

# 国内代理
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=off

WORKDIR /app

# 1️⃣ 先拷全部源码（关键！）
COPY go.mod go.sum* ./
COPY cmd/ cmd/
COPY internal/ internal/

# 2️⃣ 在容器里 tidy（此时能感知所有 import）
RUN go mod tidy

# 3️⃣ 下载依赖
RUN go mod download

# 4️⃣ 编译
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app

# ========== 运行镜像 ==========
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]