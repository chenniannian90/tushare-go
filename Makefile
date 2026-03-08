.PHONY: test build clean gen gen-specs gen-mcp gen-all verify examples fix-encoding run-examples version lint lint-fix fmt vet

# Version information
VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0-dev")
GIT_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S' 2>/dev/null || echo "unknown")

# Ldflags for version injection
LDFLAGS := -ldflags "-X main.Version=$(VERSION) -X main.GitCommit=$(GIT_COMMIT) -X main.BuildDate=$(BUILD_DATE)"

test:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

build-mcp:
	go build $(LDFLAGS) -o bin/tushare-mcp ./cmd/mcp-server

# Alias for build-mcp
build: build-mcp

build-mcp-auth:
	go build $(LDFLAGS) -o bin/mcp_server_with_auth ./cmd/mcp-server

build-gen:
	go build $(LDFLAGS) -o bin/generator ./cmd/generator

build-spec-gen:
	go build $(LDFLAGS) -o bin/spec-gen ./cmd/spec-gen

build-mcp-tools:
	@echo "MCP tools are run directly without building"

# Show version information
version:
	@echo "Version: $(VERSION)"
	@echo "Git Commit: $(GIT_COMMIT)"
	@echo "Build Date: $(BUILD_DATE)"
	@echo "Ldflags: $(LDFLAGS)"

gen: build-gen
	./bin/generator pkg/sdk/api

gen-specs: build-spec-gen
	./bin/spec-gen docs/api-directory.json internal/gen/specs

gen-mcp:
	@echo "🔨 Generating MCP tools from updated specs..."
	go run cmd/gen-mcp-tools/main.go -optimized

gen-all: gen-specs gen gen-mcp
	@echo "✅ All code generation completed successfully!"
	@echo "   - Spec files regenerated with updated descriptions"
	@echo "   - API code regenerated from updated specs"
	@echo "   - MCP tools regenerated with updated descriptions"

verify: verify-generation

verify-generation:
	@echo "🔍 Verifying code generation quality..."
	@./scripts/verify_generation.sh

examples: build-examples

build-examples:
	@echo "Building all examples..."
	@mkdir -p bin
	go build -o bin/stock_basic_example ./cmd/examples/stock_basic
	go build -o bin/daily_example ./cmd/examples/daily
	go build -o bin/daily_basic_example ./cmd/examples/daily_basic
	go build -o bin/trade_cal_example ./cmd/examples/trade_cal
	go build -o bin/financial_data_example ./cmd/examples/financial_data
	go build -o bin/index_data_example ./cmd/examples/index_data
	go build -o bin/futures_example ./cmd/examples/futures
	go build -o bin/fund_example ./cmd/examples/fund
	go build -o bin/hk_stock_example ./cmd/examples/hk_stock
	go build -o bin/boards_example ./cmd/examples/boards
	go build -o bin/sdk_usage_example ./cmd/examples/sdk_usage
	@echo "✅ Examples built successfully!"

run-examples: build-examples
	@echo "Running examples..."
	@echo "Note: Make sure TUSHARE_TOKEN environment variable is set"
	@for example in bin/*_example; do \
		if [ -x "$$example" ]; then \
			echo "\n🚀 Running $$(basename $$example)..."; \
			$$example; \
		fi; \
	done

clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

fix-encoding:
	@echo "Checking and fixing encoding issues in API spec files..."
	@echo "🔍 Scanning directory: internal/gen/specs"
	@echo ""
	@TOTAL_FILES=0; \
	CHECKED_FILES=0; \
	FIXED_FILES=0; \
	find internal/gen/specs -type f -name "*.json" -print0 | while IFS= read -r -d '' file; do \
		((TOTAL_FILES++)); \
		((CHECKED_FILES++)); \
		filename=$$(basename "$$file"); \
		if ! echo "$$filename" | iconv -f UTF-8 -t UTF-8 >/dev/null 2>&1; then \
			fixed_name=$$(echo "$$filename" | iconv -f UTF-8 -t UTF-8 -c 2>/dev/null || echo "$$filename"); \
			if [ "$$fixed_name" != "$$filename" ]; then \
				dir=$$(dirname "$$file"); \
				new_path="$$dir/$$fixed_name"; \
				mv "$$file" "$$new_path"; \
				echo "   🔧 Renamed: $$filename -> $$fixed_name"; \
				((FIXED_FILES++)); \
			fi; \
		fi; \
		python3 -c "import json, sys; exec(\
try: \
	with open('$$file', 'r', encoding='utf-8') as f: data = json.load(f); \
except UnicodeDecodeError: \
	try: \
		with open('$$file', 'r', encoding='latin-1') as f: content = f.read(); \
		with open('$$file', 'w', encoding='utf-8') as f: f.write(content); \
	except: sys.exit(1); \
except json.JSONDecodeError: \
	with open('$$file', 'r') as f: content = f.read(); \
	import re; content = re.sub(r'[\x00-\x08\x0B\x0C\x0E-\x1F\x7F]', '', content); \
	with open('$$file', 'w', encoding='utf-8') as f: f.write(content); \
)" 2>/dev/null; \
		if [ $$? -eq 0 ]; then \
			echo "   🔧 Fixed encoding in: $$filename"; \
			((FIXED_FILES++)); \
		fi; \
	done; \
	echo ""; \
	echo "📊 Summary:"; \
	echo "   Total files found: $$TOTAL_FILES"; \
	echo "   Files checked: $$CHECKED_FILES"; \
	echo "   Files fixed: $$FIXED_FILES"; \
	echo ""; \
	if [ $$FIXED_FILES -eq 0 ]; then \
		echo "✅ No encoding issues found!"; \
	else \
		echo "✅ Fixed $$FIXED_FILES file(s) with encoding issues"; \
	fi

# Run golangci-lint
lint:
	@echo "🔍 Running golangci-lint..."
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "❌ golangci-lint is not installed. Please install it first:"; \
		echo "   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
		exit 1; \
	fi
	@golangci-lint run --config .golangci.yml ./...

# Run golangci-lint with auto-fix
lint-fix:
	@echo "🔧 Running golangci-lint with auto-fix..."
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "❌ golangci-lint is not installed. Please install it first:"; \
		echo "   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
		exit 1; \
	fi
	@golangci-lint run --config .golangci.yml --fix ./...

# Format Go code
fmt:
	@echo "🎨 Formatting Go code..."
	@go fmt ./...
	@echo "✅ Code formatted successfully!"

# Run go vet
vet:
	@echo "🔍 Running go vet..."
	@go vet ./...
	@echo "✅ go vet completed successfully!"
