
# Build stage
FROM golang:1.19-rc-alpine3.15 AS builder

WORKDIR /bankapp
COPY . .

RUN go build -o main main.go
RUN apk add curl
RUN  curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run Stage
FROM alpine:latest

WORKDIR /bankapp
COPY --from=builder  /bankapp/main .
COPY --from=builder /bankapp/migrate ./migrate
COPY ./migrations/. ./migrations
# RUN sudo apt-get update -y
COPY start.sh .
#RUN sh start.sh

COPY bankapp.env .


EXPOSE 8080
CMD ["/bankapp/main"]
ENTRYPOINT ["/bankapp/start.sh"]