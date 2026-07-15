FROM golang:1.21-alpine AS builder

ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=off

WORKDIR /app

COPY go.mod go.sum* ./
COPY cmd/ cmd/
COPY internal/ internal/
COPY config/config.yaml ./config/config.yaml

RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/config.yaml .
EXPOSE 8080
CMD ["./app"]