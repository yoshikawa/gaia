# GAIA

[![CircleCI](https://circleci.com/gh/Pluslab/gaia.svg?style=svg&circle-token=f7d7b8e52887580b128e416f39b29974cb64c503)](https://circleci.com/gh/Pluslab/gaia)

Geometry App In Agriculture

## Description(詳細)

農業におけるジオメトリアプリ

オープン型圃場センシングプラットフォーム

HTTPプロトコルを用いたREST API

ベンダ独自のデータ・フォーマットに頼らず, データの再利用性を確立する.

## Requirements(必要条件)

このリポジトリを開発する時には,Dockerという仮想環境を用意します.

Homebrewが入ってるのであれば,以下のコマンドを実行

```sh
brew cask install docker
```

**Golang, React, MySQLはDockerコンテナ内にビルドされます**

| 言語/FW        | Version |
| :------------- | ------: |
| go             |  1.11.5 |
| docker         | 18.09.1 |
| docker-compose |  1.23.2 |
| node           |  11.8.0 |
| react          |  16.7.0 |
| mysql          |  8.0.14 |

## Usage(使い方)

```sh
git clone git@github.com:Pluslab/gaia.git
```

**Makefileに全てが書いています**

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

このプロジェクトは拡張性を常に保ち続ける必要があります.

## Licence(ライセンス)

[MIT](https://github.com/Pluslab/gaia/blob/master/LICENSE)

## Author(著者)

[Yoshikawa Taiki](https://github.com/yoshikawataiki)