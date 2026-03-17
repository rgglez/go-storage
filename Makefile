SHELL := /bin/bash

.PHONY: all check format vet build test generate tidy integration_test build-all latest-tags

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  check               to do static check"
	@echo "  build               to create bin directory and build"
	@echo "  generate            to generate code"
	@echo "  test                to run test"
	@echo "  build-all           to build all packages"
	@echo "  latest-tags         to show the highest git tag for each Go module"

check: vet

format:
	gofmt -w -l .

vet:
	go vet ./...

generate:
	go generate ./...
	gofmt -w -l .

build: tidy generate format check
	go build ./...

build-all:
	for f in $$(find . -name go.mod); do  \
		d=$$(dirname $$f);                \
		if [ -f "$$d/Makefile" ]; then    \
			$(MAKE) -C $$d build;         \
		else                               \
			echo "skip $$d (no Makefile)"; \
		fi;                                \
	done

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...
	go tool cover -html="coverage.txt" -o "coverage.html"

test-all:
	for f in $$(find . -name go.mod); do  \
		d=$$(dirname $$f);                \
		if [ -f "$$d/Makefile" ]; then    \
			$(MAKE) -C $$d test;          \
		else                               \
			echo "skip $$d (no Makefile)"; \
		fi;                                \
	done

tidy:
	go mod tidy
	go mod verify

tidy-all:
	for f in $$(find . -name go.mod);     \
		do make -C $$(dirname $$f) tidy;  \
	done

latest-tags:
	@for f in $$(find . -name go.mod | sort); do                              \
		d=$$(dirname $$f);                                                    \
		if [ "$$d" = "." ]; then                                              \
			rel=".";                                                          \
			pattern='^v[0-9]+\.[0-9]+\.[0-9]+$$';                            \
		else                                                                   \
			rel=$${d#./};                                                     \
			pattern="^$${rel}/v[0-9]+\.[0-9]+\.[0-9]+$$";                    \
		fi;                                                                    \
		latest=$$(git tag | grep -E "$$pattern" | sort -V | tail -1);        \
		if [ -n "$$latest" ]; then                                            \
			printf "%-45s %s\n" "$$rel" "$$latest";                          \
		else                                                                   \
			printf "%-45s %s\n" "$$rel" "(no tags)";                         \
		fi;                                                                    \
	done

clean:
	find . -type f -name 'generated.go' -delete
