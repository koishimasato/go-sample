# 概要
ドメイン駆動設計入門のサンプルコード。

# transaction
トランザクションのサンプル。

# layered
レイヤードアーキテクチャのサンプル。

## 動作確認方法
$ cd layered
$ docker-compose up -d
$ go run cmd/main.go

Index:
$ curl -i -XGET localhost:8080/users

Get:
$ curl -i -XGET localhost:8080/users/1

Post:
$ curl -XPOST localhost:8080/users/ -d '{"name": "hogehoge" }'

Put:
$ curl -XPUT localhost:8080/users/ -d '{"id": "9", "name": "hogehoge3" }'

Delete:
$ curl -i -XDELETE localhost:8080/users/1


