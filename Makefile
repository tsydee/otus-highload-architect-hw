-include .env
current_dir := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))

# Tools
export TOOLS = $(current_dir)/tools
export TOOLS_BIN = $(TOOLS)/bin
export PATH := $(TOOLS_BIN):$(PATH)

.PHONY:
run:
	docker-compose up -d --force-recreate --build apply-migration
	go run cmd/main.go

.PHONY:
lint:
	$(TOOLS_BIN)/golangci-lint run

.PHONY:
fmt-imports:
	$(TOOLS_BIN)/gofumpt -l -w .


.PHONY: install-tools
install-tools: export GOBIN=$(TOOLS_BIN)
install-tools:
	@mkdir -p $(TOOLS_BIN)
	go install github.com/pressly/goose/v3/cmd/goose@v3.10.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
	go install mvdan.cc/gofumpt@v0.6.0