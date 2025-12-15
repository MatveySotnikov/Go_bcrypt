FROM golang:1.25.1-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app/Go_bcrypt ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/Go_bcrypt /app/

CMD ["/app/Go_bcrypt"]