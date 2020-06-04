FROM golang:1.14-alpine AS build

RUN apk add --no-cache git && \
    go get -u -d github.com/golang-migrate/migrate/cmd/migrate && \
    cd $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate && \
    git checkout v4.11.0 && \
    go build -tags 'postgres' -ldflags="-X main.Version=$(git describe --tags)" -o $GOPATH/bin/migrate $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build


FROM alpine:3 AS main

RUN apk add --no-cache bash

WORKDIR /app
COPY --from=build /build/goblogs /app/goblogs
COPY --from=build /build/db/migrations /app/migrations
COPY --from=build /build/run.sh /app/run.sh
COPY --from=build /go/bin/migrate /usr/local/bin/migrate

CMD ["/app/run.sh"]