# FieldSensing

フィールドセンシング

## Description(詳細)

マルチプラットフォームフィールドセンシングシステム

## Requirements(必要条件)

このリポジトリを開発する時には,Dockerという仮想環境を用意します.

Homebrewが入ってるのであれば,以下のコマンドを実行

```sh
brew cask install docker
```

**Golang, React, MySQLはDockerコンテナ内にビルドされます**

| 言語/FW        | Version |
| :------------- | ------: |
| go             |  1.11.2 |
| docker         | 18.09.0 |
| docker-compose |  1.23.2 |
| react          |  16.6.3 |
| mysql          |  8.0.12 |

## Usage(使い方)

```sh
git clone git@github.com:Pluslab/fieldsensing.git
```

### develop backend(バックエンドの開発)

1. make docker/start
2. make api/init
3. make migrate/init
4. make migrate/up
5. make run

### develop frontend(フロントエンドの開発)

1. make docker/start
2. make npm/install
3. make npm/start

## Architecture(アーキテクチャ設計)

このプロジェクトのアーキテクチャは、クリーンアーキテクチャの原則に従います。

## Licence(ライセンス)

[MIT](https://github.com/Pluslab/FieldSensing/blob/master/LICENSE)

## Author(著者)

[Yoshikawa Taiki](https://github.com/yoshikawataiki)
