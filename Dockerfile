FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY . .
RUN go build -o .out/ufolio .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/.out/ufolio .
COPY --from=builder /app/www ./www
EXPOSE 8080
CMD ["./ufolio"]
