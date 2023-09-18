FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main/main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY ./swagger_main.yml ./swagger_main.yml

EXPOSE 3000

CMD ["/app/main"]