# MySQL

[README](/README.md) 通りに Docker コンテナに入った場合、 `#` にいると思います。  
そこから、下記のコマンドで MySQL にログインし、クエリを実行できます。

```zsh
# ログイン
mysql -u [user] -p

# ログアウト
exit
```

## 参考コマンド

参考程度に、MySQL でよく使うコマンドを数個紹介しておきます。

```sql
# データベース一覧を表示
show database;

# 指定したデータベースの中に入る
use [database];

# テーブル一覧を表示
show tables;

# 指定したテーブルのテーブル設計を表示
desc [table];

# 指定したテーブルの全レコードを表示
select * from [table];
```
