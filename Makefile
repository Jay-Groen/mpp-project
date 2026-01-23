.PHONY: test vet fmt lint coverage coverage-html coverage-gate ci

GO ?= go
TMP_DIR := .tmp
COVERPROFILE := $(TMP_DIR)/coverage.out
COVER_TXT := $(TMP_DIR)/coverage.txt
COVER_HTML := $(TMP_DIR)/coverage.html
COVERAGE_MIN ?= 0

test:
	$(GO) test ./... -count=1

vet:
	$(GO) vet ./...

fmt:
	$(GO) fmt ./...

lint:
	golangci-lint run

coverage:
	bash scripts/coverage.sh

coverage-html:
	bash scripts/coverage.sh --html

coverage-gate:
	COVERAGE_MIN=$(COVERAGE_MIN) bash scripts/coverage_gate.sh

ci: vet test lint coverage-gate
