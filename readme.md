```bash
lsof -ti :3000 | xargs kill -9

docker-compose -f docker-compose.local.yml up -d

go install github.com/swaggo/swag/cmd/swag@latest
swag init -g cmd/api/main.go --output docs
```