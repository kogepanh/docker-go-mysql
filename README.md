# Learning nuxt.js and go for docker

## 環境構築

- /docker/db/mysql/.env ファイルが必要です。
- /docker/api/go/.env ファイルが必要です。

```zsh
docker -v
docker-compose -v
```

### docker-compose

```zsh
docker-compose build
docker-compose build --no-cache
```

```zsh
docker-compose up
docker-compose up -d (バックグラウンドで実行)
```

```zsh
docker-compose down
```

```zsh
docker ps -a
docker exec -it IMAGE_ID sh
```
