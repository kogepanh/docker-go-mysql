# 必須 - 基になるイメージ
FROM golang:alpine

# 環境変数の設定
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

# 作業ディレクトリ
WORKDIR /go/src

# 実行コマンド
RUN apk update

# パッケージのインストール

# コンテナが接続用にリッスンするポートの指定
EXPOSE 8080
