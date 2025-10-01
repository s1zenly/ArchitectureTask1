FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./
RUN go build -o /temperature-api

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /temperature-api .

EXPOSE 8081
CMD ["./temperature-api"]
