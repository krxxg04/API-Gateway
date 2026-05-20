FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

ARG APP_PATH=./cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/service ${APP_PATH}

FROM alpine:3.20

RUN adduser -D -u 10001 appuser
WORKDIR /home/appuser

COPY --from=builder /bin/service ./service

USER appuser

EXPOSE 8080

CMD ["./service"]
