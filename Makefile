include .envrc

## help: print this help message
.PHONY: help
help:
	@echo 'Usage'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## run: run the web server
.PHONY: run
run:
	@go run ./cmd/web -db-dsn=${DB_DSN}