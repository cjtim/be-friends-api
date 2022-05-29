db-diff:
	go run github.com/prisma/prisma-client-go migrate diff \
	--from-url file:dev.db \
	--to-url file:dev.db

db-migrate:
	go run github.com/prisma/prisma-client-go migrate dev --name init

pre-run:
	go run github.com/prisma/prisma-client-go generate

run: pre-run
	go run main.go
