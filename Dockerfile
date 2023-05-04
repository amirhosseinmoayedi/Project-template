FROM golang:1.20-bullseye as builder

WORKDIR /app

COPY go.mod go.sum /app/

COPY . /app/

RUN CGO_ENABLED=0 go build -o url-shortner ./.

RUN chmod +x /app/main

FROM alpine:3.17.0

WORKDIR /app

COPY --from=builder /app /app

CMD ["./main"]