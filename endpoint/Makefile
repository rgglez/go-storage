SHELL := /bin/bash

.PHONY: all check format vet lint build test generate tidy integration_test release-next

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  check               to do static check"
	@echo "  build               to create bin directory and build"
	@echo "  test                to run test"
	@echo "  release-next        to tag and push the next patch version"

check: vet

format:
	@echo "go fmt"
	@go fmt ./...
	@echo "ok"

vet:
	@echo "go vet"
	@go vet ./...
	@echo "ok"

build: tidy check
	@echo "build storage"
	@go build -tags tools ./...
	@echo "ok"

test:
	@echo "run test"
	@go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...
	@go tool cover -html="coverage.txt" -o "coverage.html"
	@echo "ok"

tidy:
	@go mod tidy
	@go mod verify

release-next:
	@git_root=$$(git rev-parse --show-toplevel); \
	mod_path=$$(realpath --relative-to="$$git_root" .); \
	current=$$(git tag | grep -E "^$${mod_path}/v[0-9]+\.[0-9]+\.[0-9]+$$" | sort -V | tail -1); \
	if [ -z "$$current" ]; then echo "No tag found for $${mod_path}"; exit 1; fi; \
	version=$${current#$${mod_path}/}; \
	major=$$(echo "$$version" | cut -d. -f1 | tr -d 'v'); \
	minor=$$(echo "$$version" | cut -d. -f2); \
	patch=$$(echo "$$version" | cut -d. -f3); \
	next_tag="$${mod_path}/v$${major}.$${minor}.$$(( patch + 1 ))"; \
	echo "Tagging: $$next_tag"; \
	git tag "$$next_tag" && git push origin "$$next_tag"
