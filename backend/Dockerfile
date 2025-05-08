FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o server .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

ENV DB_USER=postgres
ENV DB_PASS=secret
ENV DB_NAME=tasks
ENV DB_HOST=db

EXPOSE 8080

CMD ["./server"]

