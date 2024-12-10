# ベースイメージとしてGoを使用
FROM golang:1.22.0-alpine

# 作業ディレクトリを設定
WORKDIR /training_manage_app_api

# Goモジュールとソースコードをコピー
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# アプリケーションを実行
CMD ["go", "run", "main.go"]