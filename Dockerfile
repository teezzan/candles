FROM golang:1.19.2-alpine3.16 AS build

WORKDIR /go/src/app

# Download git, needed for go mod download with GOPRIVATE=.
RUN apk add --update --no-cache \
    ca-certificates \
    curl \
    && rm -rf /var/cache/apk/*

RUN set -x \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Fetch await to use for Docker Compose to wait for database to be available.
RUN set -x \
    && curl -s -f -L -o await https://github.com/djui/await/releases/download/1.3.1/await-linux-amd64 \
    && chmod +x await

COPY go.mod ./
COPY go.sum ./

COPY ./cmd ./cmd
COPY ./db ./db
COPY ./internal ./internal
COPY ./scripts/startup.sh ./startup.sh

RUN CGO_ENABLED=0 go build -o /go/bin/server ./cmd/server

##

FROM alpine:3.16

WORKDIR /app

COPY --from=build /go/bin/server ./server
COPY --from=build /go/src/app/db ./db
COPY --from=build /go/src/app/startup.sh ./startup.sh
COPY --from=build /go/src/app/migrate ./migrate
COPY --from=build /go/src/app/await ./await

EXPOSE 8090

ENTRYPOINT ["./startup.sh", "./server"]
