FROM golang:alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

ARG run=graphQL


RUN CGO_ENABLED=0 GOOS=linux go build -o ./api ./cmd/$run


FROM alpine:latest

WORKDIR /app

COPY --from=build /app/api ./api

EXPOSE 8080
ENTRYPOINT  ["./api"]




