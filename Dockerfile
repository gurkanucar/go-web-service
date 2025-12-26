    FROM golang:1.24-alpine AS backend-builder
    WORKDIR /app
    RUN apk add --no-cache git gcc musl-dev
    COPY go.mod go.sum ./
    RUN go mod download
    COPY . .
    RUN go install github.com/swaggo/swag/cmd/swag@latest
    RUN swag init -g cmd/api/main.go --output docs
    RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o project cmd/api/main.go
    
    FROM alpine:latest
    WORKDIR /app
    RUN apk add --no-cache ca-certificates tzdata libc6-compat
    COPY --from=backend-builder /app/project .
    COPY --from=backend-builder /app/docs ./docs
    EXPOSE 8080
    CMD ["./project"]