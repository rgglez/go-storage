SHELL := /bin/bash

.PHONY: all check format lint build test generate tidy release-next

-include Makefile.env

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  check               to do static check"
	@echo "  build               to create bin directory and build"
	@echo "  generate            to generate code"
	@echo "  test                to run test"
	@echo "  release-next        to tag and push the next patch version"

check: vet

format:
	go fmt ./...

vet:
	go vet ./...

generate:
	@echo "generate code"
	go generate ./...
	go fmt ./...

build: tidy generate check
	go build ./...

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic -v .
	go tool cover -html="coverage.txt" -o "coverage.html"

integration_test:
	go test -race -count=1 -covermode=atomic -v ./tests

tidy:
	go mod tidy
	go mod verify

clean:
	@echo "clean generated files"
	find . -type f -name 'generated.go' -delete

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
