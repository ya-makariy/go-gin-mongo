FROM golang:1.17-alpine AS build

WORKDIR /tmp/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./go-gin-mongo

FROM alpine:3.9

COPY --from=build /tmp/app/go-gin-mongo /app/go-gin-mongo

EXPOSE 8080

CMD [ "/app/go-gin-mongo" ]