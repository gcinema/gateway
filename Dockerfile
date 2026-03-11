FROM golang:1.26.0 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/gateway ./cmd

FROM alpine as runner

WORKDIR /app
COPY --from=builder /app/bin/gateway /app

EXPOSE 8080

CMD ["/app/gateway"]