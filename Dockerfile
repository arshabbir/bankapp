FROM golang:1.19-rc-alpine3.15 AS builder

WORKDIR /bankapp
COPY . .

RUN go build -o main main.go


FROM alpine:latest

WORKDIR /bankapp
COPY --from=builder /bankapp/main .
COPY bankapp.env .


EXPOSE 8080
ENTRYPOINT ["/bankapp/main"]