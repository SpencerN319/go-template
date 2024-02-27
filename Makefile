# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_\/0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-10s ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.PHONY: build
## Build local environment
build:
	@docker-compose build

.PHONY: up
## Run local environment
up:
	@docker-compose up --wait --build

.PHONY: watch
## Watch for file changes and rebuild image
watch:
	@docker-compose watch --no-up

.PHONY: down
## Stop local environment
down:
	@docker-compose down -v

.PHONY: clean
## Remove dangling docker images (i.e. untagged "<none>" images)
clean:
	@go clean -testcache
	@$(shell rm -rf target/*)
	@docker rmi $(shell docker images -f "dangling=true" -q)

.PHONY: integration-test
## Run local integration tests
integration-test:
	@echo '${GREEN}Integration Tests${RESET}'
	@go clean -testcache
	@go test -race --tags=integration -timeout 30s -v -coverprofile integration_coverage.out ./...
	@go tool cover -html=integration_coverage.out
	@rm integration_coverage.out

.PHONY: unit-test
## Run unit tests & store coverage log, Server and Client coverage generated separately
unit-test:
	@echo '${GREEN}Unit Tests${RESET}'
	@go clean -testcache
	@go test -race --tags=unit -timeout 30s -coverprofile unit_coverage.out ./...
	@go tool cover -html=unit_coverage.out
	@rm unit_coverage.out
