FROM golang:1.27.2-alpine3.24 AS builder
WORKDIR /app
COPY go.mod main.go ./
RUN go build -o main .

FROM alpine:3.24.2
WORKDIR /root
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
