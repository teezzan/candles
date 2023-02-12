# Candles

## The Ultimate OHLC API.

![Logo](./docs/logo.png)

> **Candles** has the following features

- [x] Upload OHLC file in CSV format to be processed and stored.
- [x] Granulated fetch based on Symbol and time range.
- [x] Highly scalable process pipeline of very large CSV files using AWS S3 and SQS queues.
- [x] Easy to setup
- [x] Documentation with [OpenAPI](https://leaguey.herokuapp.com/swagger-html) (https://leaguey.herokuapp.com/swagger-html).
- [ ] Caching of results for efficiency

**Note**: **I have a deployed version of the aplication found [here on Heroku](https://leaguey.herokuapp.com/swagger-html)**

## Images

<details>
<summary>Here are samples.</summary>

![docs](./images/openAPI.png)

</details>

## Built With

- [Go](https://go.dev/)
- [SQLx](https://github.com/jmoiron/sqlx)
- [Gin](https://github.com/gin-gonic/gin)
- [MySQL](https://www.mysql.com/)
- [AWS SDK Go V2](https://github.com/aws/aws-sdk-go-v2)

## Getting Started

This application can be deployed locally by the following ways.

- Using `docker-compose`
- Using the local environment.

### Prerequisites

You require a `.env` file to get the application running on your local computer. A sample `env` file is as follows.

```
## Database Configuration
DB_HOST=localhost
DB_NAME=candles
DB_USERNAME=root
DB_PASSWORD=mysqltest
DB_PORT=3306


## Server Configuration
SERVER_PORT=8090

## AWS Configurations
AWS_ACCESS_KEY_ID=AKIAJ4Z7Z7Z7Z7Z7Z7Z7
AWS_SECRET_ACCESS_KEY=pw777yyh777h66+ff777777777777777777777777

```

You can make a copy of the provided `.env.example` file, rename it as `.env` and populate it with the appropriate vales. Ensure that the `.env` file contains the necessary credentials. The system will not start if important credentials like `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` are not present. Other credentials have a reasonable default set in the app. These can found in `internal/config/defaults.go`.

You may also need to have Docker installed if you wish to follow that path.

### Docker-compose

You need to run the project by running.

```bash
make docker-up
```

Ensure that the `.env` file contains the necessary credentials.

### Local Environment

Foremost, you have to clone this repository and install all the dependencies with the following commands.

```bash
git clone https://github.com/teezzan/candles.git
cd candles
go get ./...
make deps
make migrations-up
```

**Note**: `DB_NAME` will not be created automatically and would need to ba manually created.

Run `make serve` to start the local server.

### Make Commands

run `make help` to see a list of supported commands.

```bash
$ make help
    ------------------------------------------------------------------------
    candles
    ------------------------------------------------------------------------
    serve                          Run locally
    build                          Build application binaries
    deps                           Install build dependencies
    deps-moq                       Install build dependencies: moq
    deps-godotenv                  Install build dependencies: godotenv
    deps-swag                      Install build dependencies: Swag
    image                          Create Docker image
    docker-only-up                 Start Docker image
    docker-up                      Start docker-compose
    docker-down                    Stop docker-compose
    test                           Run unit tests
    docs                           Generates OpenAPI docs using https://github.com/swaggo/swag
```

## API Reference/Documentation

A swagger documentation can be found at `http://localhost:3000/swagger-html`

## Authors

**[Taiwo Yusuf](https://github.com/teezzan/)**
