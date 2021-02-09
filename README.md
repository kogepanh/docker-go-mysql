# Golang and MySQL for Docker Template

## Docker 環境構築

本来、環境変数等は .gitignore で ignore すべきですが、テンプレートであることと勉強用なので重要なパラメータは存在していないため公開しています。

> 本来は以下のディレクトリに `.env` ファイルが必要です。
>
> - /docker/db/mysql/.env
> - /docker/api/go/.env

```zsh
# docker/docker-composeが使えるか確認
docker -v
docker-compose -v
```

```zsh
# docker-compose.ymlで定義したserviceのビルドを実行
docker-compose build

# キャッシュのせいで設定変更が反映されない時はオプションをつけるべし
# docker-compose build --no-cache
```

```zsh
# コンテナを起動
docker-compose up

# -d でバックグラウンド実行できるが、初学者はやめとくべし
# docker-compose up -d
```

```zsh
# コンテナ一覧を表示
# 入りたいコンテナのCONTAINER IDをコピーする
docker ps -a
```

```zsh
# コンテナの中に入る
docker exec -it [IMAGE_ID] sh
```

```zsh
# docker-compose.ymlで定義したserviceを参考に、コンテナを停止
# docker-compose.ymlのあるディレクトリで実行する必要がある
docker-compose down

# --rmi all でイメージを削除できる
# docker-compose down --rmi all
```

```zsh
# 全コンテナ削除
docker rm $(docker ps -q -a)

# 指定コンテナ削除
docker rm [CONTAINER_ID]
```

```zsh
# 全イメージ削除
docker rmi $(docker images -q)

# 指定イメージ削除
docker rmi [IMAGE_ID]
```

## Docker コンテナの利用方法

- [MySQL に関するドキュメント](/docker/db/README.md)
- [Golang に関するドキュメント](/docker/api/README.md)
