# FieldSensing

[![CircleCI](https://circleci.com/gh/Pluslab/fieldsensing.svg?style=svg&circle-token=f7d7b8e52887580b128e416f39b29974cb64c503)](https://circleci.com/gh/Pluslab/fieldsensing)

FieldSensing

## Description

Multi-platform field sensing system

## Requirements

| 言語/FW        | Version |
| :------------- | ------: |
| go             |  1.11.2 |
| docker         | 18.09.0 |
| docker-compose |  1.23.2 |
| react          |  16.6.3 |
| mysql          |  8.0.12 |

## Usage

```sh
git clone git@github.com:Pluslab/fieldsensing.git
```

### develop backend

1. make docker/start
2. make api/init
3. make migrate/init
4. make migrate/up
5. make run

access `localhost:8080`

### develop frontend

1. make docker/start
2. make npm/install
3. make npm/start

access `localhost:8081`

## Architecture

The architecture of the project follows the principles of Clean Architecture.

## Licence

[MIT](https://github.com/Pluslab/FieldSensing/blob/master/LICENSE)

## Author

[Yoshikawa Taiki](https://github.com/yoshikawataiki)
