FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o .out/ufolio .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/.out/ufolio .
RUN mkdir -p /root/.ufolio
EXPOSE 3000
CMD ["./ufolio"]
