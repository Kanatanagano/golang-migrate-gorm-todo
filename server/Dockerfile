# ベースイメージ
FROM golang:1.23.2 AS base
ENV TZ=Asia/Tokyo
WORKDIR /src/server

# ローカル開発用
FROM base AS local

RUN go install -v github.com/air-verse/air@latest
RUN go install -v github.com/gin-gonic/gin@latest

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y \
    jq

# アプリをビルド
FROM base AS builder
RUN --mount=type=bind,target=. GOOS=linux GOARCH=amd64 go build -o /go/bin/app


# 本番用
FROM base AS production
COPY --from=builder /go/bin/app /app
CMD ["/app"]
