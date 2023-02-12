# Candles - The Ultimate OHLC API

![Logo](./docs/logo.png)

Candles is a highly scalable OHLC API with the following features:

- Upload OHLC file in CSV format to be processed and stored
- Granulated fetch based on Symbol and time range
- Process pipeline for very large CSV files using AWS S3 and SQS queues
- Easy setup
- Documentation with OpenAPI (https://candles.onrender.com/docs/)

## Deployment

This application can be deployed locally in two ways:

- Using `docker-compose`
- Using the local environment

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

To run the application on your local computer, you need a .env file containing the following:

```
## Database Configuration
DB_HOST=localhost
DB_NAME=ohlc
DB_USERNAME=root
DB_PASSWORD=mysqltest
DB_PORT=3306


## Server Configuration
SERVER_HOST=localhost
SERVER_PORT=8090

## AWS Configuration
AWS_ACCESS_KEY_ID=AKIAJ4Z7Z7Z7Z7Z7Z7Z7
AWS_SECRET_ACCESS_KEY=pw777yyh777h66+ff777777777777777777777777

# AWS S3 Bucket Configuration
S3_REGION=
S3_BUCKET=
S3_PRESIGN_URL_EXPIRY_TIME=

# AWS SQS Configuration
SQS_REGION=
SQS_QUEUE=

# Cron Job Configuration
CRON_JOB_FREQUENCY_IN_MINUTES=

```

A sample .env file named `.env.example` is provided. You can make a copy, rename it to `.env`, and update the values. The system will not start if important credentials, such as `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`, are missing. The default values for other credentials can be found in `internal/config/defaults.go`.

You may also need to have Docker installed to use the docker-compose method.

### Docker-compose

You need to run the project by running.

```bash
make docker-up
```

Ensure that the `.env` file contains the necessary credentials.

### Local Environment

To deploy the application locally, follow these steps:

- Clone the repository: `git clone https://github.com/teezzan/candles.git`
- Navigate to the project directory: `cd candles`
- Install dependencies: `go get ./... && make deps`
- Create the database: `make migrations-up`
- Start the local server: `make serve`

Note: The database name specified in DB_NAME must be created manually.

To view the available make commands, run make help.

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
    deps-migrate                   Install build dependencies: Migrate
    image                          Create Docker image
    docker-only-up                 Start Docker image
    docker-up                      Start docker-compose
    docker-down                    Stop docker-compose
    test                           Run unit tests
    docs                           Generates OpenAPI docs using https://github.com/swaggo/swag
    migrations-up                  Run migrations
```

## API Reference/Documentation

The API reference and documentation are available in swagger format on [Render](https://candles.onrender.com/docs/) or locally at `http://localhost:3000/docs`

![docs](./docs/openAPI.png)

## Authors

**[Taiwo Yusuf](https://github.com/teezzan/)**
