PROJECT_NAME=withdrawdeposit

# DATABASE
DB_USER=root
DB_PASSWORD=root
DB_HOST=localhost
DB_PORT=3306
DB_NAME=disbursement

install:
	cd .. && go get -u github.com/swaggo/swag/cmd/swag && go get -u github.com/cosmtrek/air && cd ${PROJECT_NAME} && swag init

local:
	air -c config/.air.toml

test:
	go test -v -cover ./...
test-cover:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out ; rm -f coverage.out

migrate-up:
	migrate -source file:./scripts/migrations/ -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable up

migrate-down:
	migrate -source file:./scripts/migrations/ -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable down

dev-up:
	docker-compose -f docker-compose-dev.yml up -d --build

dev-down:
	docker-compose -f docker-compose-dev.yml down

prod-up:
	docker-compose up -d --build

prod-down:
	docker-compose down