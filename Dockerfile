# デプロイ用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.23.1-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# ----------------------------------------

# デプロイ用のコンテナ
FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# ----------------------------------------

# ローカル開発環境で利用するホットリロード環境
FROM golang:1.23.1 as dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air"]
