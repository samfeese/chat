FROM golang:1.24.3-bookworm AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o chat-api ./cmd/api

# Final Stage
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/chat-api .
RUN chmod +x /app/chat-api

EXPOSE 8080
ENTRYPOINT ["/app/chat-api"]
