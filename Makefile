.PHONY: lint fmt vet test clean help

# 默认目标
help:
	@echo "可用命令:"
	@echo "  make lint    - 运行代码检查 (golangci-lint)"
	@echo "  make fmt     - 格式化代码"
	@echo "  make vet     - 运行 go vet"
	@echo "  make test    - 运行测试"
	@echo "  make clean   - 清理构建文件"

# 代码检查 - 使用 golangci-lint
lint:
	@echo "Running golangci-lint..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not found, installing..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin latest; \
		golangci-lint run ./...; \
	fi

# 备用 lint - 使用 go vet + gofmt (如果没有 golangci-lint)
lint-simple:
	@echo "Running go fmt check..."
	@test -z $$(gofmt -l .) || (echo "代码格式错误，请运行 'make fmt'" && exit 1)
	@echo "Running go vet..."
	go vet ./...

# 格式化代码
fmt:
	@echo "Formatting code..."
	gofmt -s -w .

# 运行 go vet
vet:
	@echo "Running go vet..."
	go vet ./...

# 运行测试
test:
	@echo "Running tests..."
	go test -v ./...

# 清理
clean:
	@echo "Cleaning..."
	go clean
