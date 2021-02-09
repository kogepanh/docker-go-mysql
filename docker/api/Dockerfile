# 必須 - 基になるイメージ
FROM golang:alpine

# 環境変数の設定
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

# 作業ディレクトリ
WORKDIR /go/src
# ホットリロード環境
# WORKDIR /go/src/hot_reload_docker

# 実行コマンド
RUN apk update

# コンテナが接続用にリッスンするポートの指定
EXPOSE 8080
