FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/AdventureAI .

# ---

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/AdventureAI .
COPY games games/

# Expose port 8080 to the outside world
EXPOSE 8080

CMD ["./AdventureAI", "-file", "games/905.z5", "-mcp"]
