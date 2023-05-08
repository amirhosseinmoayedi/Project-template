MAIN_PACKAGE_PATH := ./.
BINARY_NAME := app
DOCKER_IMAGE_NAME := app
DOCKER_IMAGE_VERSION := 1.0
DATABASE_DSN := "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

.PHONY: confirm
confirm:
	@echo -n "Are you sure? [y/N]"  && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	@git diff --exit-code


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	@go fmt ./...
	@go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	@go mod verify #  verifying the dependencies
	@go vet ./... #  static analysis
	@go run honnef.co/@go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./... #  static analysis
	@go run @golang.org/x/vuln/cmd/@govulncheck@latest ./... # checking for vulnerabilities
	@go test -race -buildvcs -vet=off ./...


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	@go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	@go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	@go tool cover -html=/tmp/coverage.out

## build: build the application
.PHONY: build
build:
    # Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
	@go build -o=/tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: run the  application
.PHONY: run
run: build
	/tmp/bin/${BINARY_NAME}

## run/live: run the application with reloading on file changes
.PHONY: run/live
run/live:
	@air serve


# ==================================================================================== #
# Docker
# ==================================================================================== #
# build docker image
.PHONY: build/image
build/image:
	@docker build -t ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_NAME} .

.PHONY: compose/up
compose/up:
	@docker compose up

.PHONY: compose/up-detach
compose/up:
	@docker compose up -d

.PHONY: compose/rebuild
compose/up:
	@docker compose up -d --build --force-recreate

.PHONY: compose/teardown
compose/up:
	@docker compose down -v


# ==================================================================================== #
# DB Migration
# ==================================================================================== #
.PHONY: migration/up
migration_up:
	@migrate -path internall/infrastructure/migration/ -database ${DATABASE_DSN} -verbose up

.PHONY: migration/down
migration_down:
	@migrate -path internall/infrastructure/migration/ -database ${DATABASE_DSN} -verbose down

.PHONY: migration/drop
dropdb:
	@docker exec -it db dropdb sample_db

.PHONY: migration/create
	@docker exec -it db createdb --username=root --owner=root sample_db