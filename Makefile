SHELL := /bin/bash

SERVICE ?=

.PHONY: all check format vet build test generate tidy integration_test build-all latest-tags \
	next-tag-service push-tag-service next-tag-credential push-tag-credential \
	next-tag-endpoint push-tag-endpoint next-tag push-next-tag

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  check               to do static check"
	@echo "  build               to create bin directory and build"
	@echo "  generate            to generate code"
	@echo "  test                to run test"
	@echo "  build-all           to build all packages"
	@echo "  latest-tags         to show the highest git tag for each Go module"
	@echo "  next-tag-service    create next PATCH tag for services/SERVICE (e.g. make next-tag-service SERVICE=oss)"
	@echo "  push-tag-service    push latest services/SERVICE tag to origin"
	@echo "  next-tag-credential create next PATCH tag for credential"
	@echo "  push-tag-credential push latest credential tag to origin"
	@echo "  next-tag-endpoint   create next PATCH tag for endpoint"
	@echo "  push-tag-endpoint   push latest endpoint tag to origin"
	@echo "  next-tag            create next PATCH tag for root module"
	@echo "  push-next-tag       push latest root tag to origin"

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

next-tag-service:
	@if [ -z "$(SERVICE)" ]; then echo "Usage: make next-tag-service SERVICE=<name>"; exit 1; fi
	@latest=$$(git tag | grep -E "^services/$(SERVICE)/v[0-9]+\.[0-9]+\.[0-9]+$$" | sort -V | tail -1); \
	if [ -z "$$latest" ]; then echo "No tags found for services/$(SERVICE)"; exit 1; fi; \
	version=$$(echo "$$latest" | grep -oE 'v[0-9]+\.[0-9]+\.[0-9]+'); \
	major=$$(echo "$$version" | cut -d. -f1 | tr -d 'v'); \
	minor=$$(echo "$$version" | cut -d. -f2); \
	patch=$$(echo "$$version" | cut -d. -f3); \
	new_tag="services/$(SERVICE)/v$${major}.$${minor}.$$((patch + 1))"; \
	echo "Creating tag: $$new_tag"; \
	git tag "$$new_tag"

push-tag-service:
	@if [ -z "$(SERVICE)" ]; then echo "Usage: make push-tag-service SERVICE=<name>"; exit 1; fi
	@latest=$$(git tag | grep -E "^services/$(SERVICE)/v[0-9]+\.[0-9]+\.[0-9]+$$" | sort -V | tail -1); \
	if [ -z "$$latest" ]; then echo "No tags found for services/$(SERVICE)"; exit 1; fi; \
	echo "Pushing tag: $$latest"; \
	git push origin "$$latest"

next-tag-credential:
	@latest=$$(git tag | grep -E "^credential/v[0-9]+\.[0-9]+\.[0-9]+$$" | sort -V | tail -1); \
	if [ -z "$$latest" ]; then echo "No credential tags found"; exit 1; fi; \
	version=$$(echo "$$latest" | grep -oE 'v[0-9]+\.[0-9]+\.[0-9]+'); \
	major=$$(echo "$$version" | cut -d. -f1 | tr -d 'v'); \
	minor=$$(echo "$$version" | cut -d. -f2); \
	patch=$$(echo "$$version" | cut -d. -f3); \
	new_tag="credential/v$${major}.$${minor}.$$((patch + 1))"; \
	echo "Creating tag: $$new_tag"; \
	git tag "$$new_tag"

push-tag-credential:
	@latest=$$(git tag | grep -E "^credential/v[0-9]+\.[0-9]+\.[0-9]+$$" | sort -V | tail -1); \
	if [ -z "$$latest" ]; then echo "No credential tags found"; exit 1; fi; \
	echo "Pushing tag: $$latest"; \
	git push origin "$$latest"

next-tag-endpoint:
	@latest=$$(git tag | grep -E "^endpoint/v[0-9]+\.[0-9]+\.[0-9]+$$" | sort -V | tail -1); \
	if [ -z "$$latest" ]; then echo "No endpoint tags found"; exit 1; fi; \
	version=$$(echo "$$latest" | grep -oE 'v[0-9]+\.[0-9]+\.[0-9]+'); \
	major=$$(echo "$$version" | cut -d. -f1 | tr -d 'v'); \
	minor=$$(echo "$$version" | cut -d. -f2); \
	patch=$$(echo "$$version" | cut -d. -f3); \
	new_tag="endpoint/v$${major}.$${minor}.$$((patch + 1))"; \
	echo "Creating tag: $$new_tag"; \
	git tag "$$new_tag"

push-tag-endpoint:
	@latest=$$(git tag | grep -E "^endpoint/v[0-9]+\.[0-9]+\.[0-9]+$$" | sort -V | tail -1); \
	if [ -z "$$latest" ]; then echo "No endpoint tags found"; exit 1; fi; \
	echo "Pushing tag: $$latest"; \
	git push origin "$$latest"

next-tag:
	@latest=$$(git tag | grep -E "^v[0-9]+\.[0-9]+\.[0-9]+$$" | sort -V | tail -1); \
	if [ -z "$$latest" ]; then echo "No root tags found"; exit 1; fi; \
	major=$$(echo "$$latest" | cut -d. -f1 | tr -d 'v'); \
	minor=$$(echo "$$latest" | cut -d. -f2); \
	patch=$$(echo "$$latest" | cut -d. -f3); \
	new_tag="v$${major}.$${minor}.$$((patch + 1))"; \
	echo "Creating tag: $$new_tag"; \
	git tag "$$new_tag"

push-next-tag:
	@latest=$$(git tag | grep -E "^v[0-9]+\.[0-9]+\.[0-9]+$$" | sort -V | tail -1); \
	if [ -z "$$latest" ]; then echo "No root tags found"; exit 1; fi; \
	echo "Pushing tag: $$latest"; \
	git push origin "$$latest"

clean:
	find . -type f -name 'generated.go' -delete
