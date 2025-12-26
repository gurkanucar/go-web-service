lsof -ti :3000 | xargs kill -9

docker-compose -f docker-compose.local.yml up -d