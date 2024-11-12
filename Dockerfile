FROM golang:1.23.1-alpine3.20 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /server cmd/main.go

FROM alpine:3.20.3

RUN mkdir -p /app/logs

COPY --from=build /server /server

ENTRYPOINT ["/server"]
