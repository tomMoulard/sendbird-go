SHELL := /bin/bash

.DEFAULT_GOAL := all
.PHONY: all
all: ## build pipeline
all: mod inst gen build spell lint test

.PHONY: ci
ci: ## CI build pipeline
# ci: all check mocktail diff
ci: all check diff # for now mocktail seems to be broken

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## remove files created during build pipeline
	$(call print-target)
	rm -rf dist
	rm -f coverage.*
	rm -f '"$(shell go env GOCACHE)/../golangci-lint"'
	go clean -i -cache -testcache -modcache -fuzzcache -x

.PHONY: mod
mod: ## go mod tidy
	$(call print-target)
	go mod tidy -x

.PHONY: inst
inst: ## go install tools
	$(call print-target)
	pip install --user yamllint

.PHONY: gen
gen: ## go generate
	$(call print-target)
	go generate ./...

.PHONY: mocktail
mocktail: ## go mocktail
	$(call print-target)
	go tool mocktail -e ./...

.PHONY: build
build: ## goreleaser build
build:
	$(call print-target)
	go tool goreleaser build --clean --single-target --snapshot

.PHONY: spell
spell: ## misspell
	$(call print-target)
	go tool misspell -error -locale=US -w **.md

.PHONY: lint
lint: ## golangci-lint
	$(call print-target)
	yamllint .
	go tool goreleaser check
	go tool golangci-lint run

.PHONY: check
check: ## govulncheck
	$(call print-target)
	go tool govulncheck -scan package ./...
	go tool govulncheck ./...

.PHONY: test
test: ## go test
	$(call print-target)
	go test -race -covermode=atomic -coverprofile=coverage.out -coverpkg=./... ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: diff
diff: ## git diff
	$(call print-target)
	git diff --exit-code
	RES=$$(git status --porcelain) ; if [ -n "$$RES" ]; then echo $$RES && exit 1 ; fi


define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef
