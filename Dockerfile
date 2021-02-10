FROM golang:1.15.2-alpine3.12

ENV GO111MODULE=on

WORKDIR /app

COPY . ./

# Download project dependencies.

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

RUN apk add --no-cache ca-certificates \
    dpkg \
    gcc \
    git \
    musl-dev \
    bash

CMD [ "/app/main" ]