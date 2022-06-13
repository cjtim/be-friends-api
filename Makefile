REDIS_PORT=3379
REDIS_CONTAINER=be-friends-api-redis

POSTGRES_PORT=6432
POSTGRES_CONTAINER=be-friends-api-db

# Secret
LINE_CLIENT_ID=
LINE_SECRET_ID=
DATABASE_URL=postgresql://postgres:postgres@localhost:$(POSTGRES_PORT)/be-friends?sslmode=disable
BACKET_NAME=
GCLOUD_CREDENTIAL=

pre-run:
	docker run --rm -d --name $(REDIS_CONTAINER) -p $(REDIS_PORT):6379 redis:7-alpine || true

	docker run --rm -d --name $(POSTGRES_CONTAINER) -p $(POSTGRES_PORT):5432 \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_PASSWORD=postgres \
	-e POSTGRES_DB=be-friends \
	-v $(PWD)/tools/db:/docker-entrypoint-initdb.d \
	postgres:14.3-alpine || true

run: pre-run
	REDIS_URL=localhost:$(REDIS_PORT) \
	DATABASE_URL=$(DATABASE_URL) \
	LINE_CLIENT_ID=$(LINE_CLIENT_ID) \
	LINE_SECRET_ID=$(LINE_SECRET_ID) \
	BACKET_NAME=$(BACKET_NAME) \
	GCLOUD_CREDENTIAL=$(GCLOUD_CREDENTIAL) \
	go run main.go

clean:
	docker rm -f $(REDIS_CONTAINER) || true
	docker rm -f $(POSTGRES_CONTAINER) || true
