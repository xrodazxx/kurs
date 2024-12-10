FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/
RUN go build -o main .  

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache libc6-compat  

COPY --from=builder /app/cmd/main .

ENTRYPOINT ["./main"]