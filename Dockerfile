FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o .out/ufolio .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/.out/ufolio .
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates
EXPOSE 8080
CMD ["./ufolio"]
