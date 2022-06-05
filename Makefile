REDIS_PORT=3379
REDIS_CONTAINER=be-friends-api-redis

db-diff:
	go run github.com/prisma/prisma-client-go migrate diff \
	--from-url file:dev.db \
	--to-url file:dev.db

db-migrate:
	go run github.com/prisma/prisma-client-go migrate dev --name init

pre-run:
	docker run --rm -d --name $(REDIS_CONTAINER) -p $(REDIS_PORT):6379 redis:7-alpine || true

run: pre-run
	REDIS_URL=localhost:$(REDIS_PORT) go run main.go

clean:
	docker rm -f $(REDIS_CONTAINER) || true