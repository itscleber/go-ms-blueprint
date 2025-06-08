FROM golang:1.23-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go

FROM golang:1.23-alpine AS dev
RUN apk add --no-cache git curl
WORKDIR /app
EXPOSE 8080
CMD ["/bin/sh"]

FROM alpine:3.18 as production
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
