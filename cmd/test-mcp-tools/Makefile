.PHONY: test build clean help

# 默认目标
.DEFAULT_GOAL := help

# 编译测试程序
build:
	@echo "编译测试程序..."
	@go build -o bin/test-mcp-tools cmd/test-mcp-tools/main.go
	@echo "✅ 编译完成: bin/test-mcp-tools"

# 运行测试
test: build
	@echo "运行测试..."
	@./bin/test-mcp-tools

# 直接运行（不编译）
run:
	@echo "运行测试程序..."
	@go run cmd/test-mcp-tools/main.go

# 清理
clean:
	@echo "清理编译文件..."
	@rm -f bin/test-mcp-tools
	@echo "✅ 清理完成"

# 检查环境
check:
	@echo "检查环境..."
	@if [ -z "$(TUSHARE_TOKEN)" ]; then \
		echo "❌ TUSHARE_TOKEN 未设置"; \
		echo "请设置: export TUSHARE_TOKEN=你的token"; \
		exit 1; \
	else \
		echo "✅ TUSHARE_TOKEN 已设置"; \
	fi

# 快速测试（只测试几个工具）
quick:
	@echo "快速测试（仅测试部分工具）..."
	@go run cmd/test-mcp-tools/main.go

# 生成报告
report:
	@echo "生成测试报告..."
	@if [ -f "tests/mcp-tools-test-report.md" ]; then \
		cat tests/mcp-tools-test-report.md; \
	else \
		echo "❌ 报告文件不存在，请先运行测试"; \
	fi

# 帮助
help:
	@echo "MCP 工具测试程序 - 可用命令:"
	@echo ""
	@echo "  make build       - 编译测试程序"
	@echo "  make test        - 编译并运行测试"
	@echo "  make run         - 直接运行测试（不编译）"
	@echo "  make quick       - 快速测试"
	@echo "  make check       - 检查环境配置"
	@echo "  make report      - 查看测试报告"
	@echo "  make clean       - 清理编译文件"
	@echo "  make help        - 显示此帮助信息"
	@echo ""
	@echo "使用示例:"
	@echo "  export TUSHARE_TOKEN=你的token"
	@echo "  make test"
	@echo ""
