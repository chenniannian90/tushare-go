package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"tushare-go/pkg/mcp/server"
)

func main() {
	fmt.Println("=== Tushare MCP HTTP 服务器示例 ===\n")

	// 配置环境变量
	os.Setenv("TUSHARE_TOKEN", "your-tushare-token")
	os.Setenv("MCP_API_KEY", "your-mcp-api-key")
	os.Setenv("MCP_REQUIRE_AUTH", "true")
	os.Setenv("LOG_LEVEL", "INFO")

	// 创建 SDK 客户端
	config := server.DefaultServerConfig()
	client, err := config.CreateSDKClient()
	if err != nil {
		log.Fatalf("创建 SDK 客户端失败: %v", err)
	}

	// 创建 HTTP MCP 服务器
	mcpServer := server.NewHTTPMCPServer(client, config)

	// 启动服务器
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 在 goroutine 中启动服务器
	go func() {
		addr := ":8080"
		if err := mcpServer.Start(ctx, addr); err != nil && err != context.Canceled {
			log.Printf("服务器错误: %v", err)
		}
	}()

	fmt.Println("✅ 服务器正在启动...")
	time.Sleep(1 * time.Second) // 等待服务器启动

	fmt.Println("\n📋 服务器功能:")
	fmt.Println("  1. MCP 协议端点:")
	fmt.Println("     POST /mcp - MCP 协议请求")
	fmt.Println()
	fmt.Println("  2. RESTful API 端点:")
	fmt.Println("     POST /api/v1/{module}/{tool} - 直接调用工具")
	fmt.Println("     POST /api/v1/hk_stock?tool=hk_basic - 统一端点")
	fmt.Println()
	fmt.Println("  3. 管理端点:")
	fmt.Println("     GET  /health - 健康检查")
	fmt.Println("     GET  /metrics - 服务指标")

	fmt.Println("\n🔧 API 调用示例:")

	// MCP 协议调用示例
	fmt.Println("\n1️⃣ MCP 协议方式:")
	fmt.Println("   curl -X POST http://localhost:8080/mcp \\")
	fmt.Println(`     -H "Content-Type: application/json" \`)
	fmt.Println(`     -d '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"clientInfo":{"name":"test-client","version":"1.0.0","apiKey":"your-mcp-api-key"}}}'`)

	// RESTful API 调用示例
	fmt.Println("\n2️⃣ RESTful API 方式:")
	fmt.Println("   curl -X POST http://localhost:8080/api/v1/hk_stock/hk_basic \\")
	fmt.Println(`     -H "Content-Type: application/json" \`)
	fmt.Println(`     -d '{"ts_code": "00700.HK", "list_date": "20240101"}'`)

	// 统一端点调用示例
	fmt.Println("\n3️⃣ 统一端点方式:")
	fmt.Println("   curl -X POST 'http://localhost:8080/api/v1/hk_stock?tool=hk_basic' \\")
	fmt.Println(`     -H "Content-Type: application/json" \`)
	fmt.Println(`     -d '{"ts_code": "00700.HK", "list_date": "20240101"}'`)

	// 健康检查示例
	fmt.Println("\n4️⃣ 健康检查:")
	fmt.Println("   curl http://localhost:8080/health")

	// 指标查询示例
	fmt.Println("\n5️⃣ 服务指标:")
	fmt.Println("   curl http://localhost:8080/metrics")

	fmt.Println("\n📊 服务器状态:")
	fmt.Println("  地址: http://localhost:8080")
	fmt.Println("  协议: HTTP + JSON-RPC 2.0")
	fmt.Println("  工具数量: 195")
	fmt.Println("  认证: 启用 (API Key)")

	fmt.Println("\n🔒 认证方式:")
	fmt.Println("  1. HTTP Header:")
	fmt.Println("     Authorization: Bearer your-mcp-api-key")
	fmt.Println("  2. 请求参数:")
	fmt.Println(`     {"clientInfo": {"apiKey": "your-mcp-api-key"}}`)

	fmt.Println("\n⚡ 支持的模块:")
	modules := []string{
		"hk_stock", "bond", "etf", "forex", "fund", "futures",
		"index", "llm_corpus", "options", "spot", "us_stock",
		"stock_basic", "stock_board", "stock_feature", "stock_financial",
		"stock_fund_flow", "stock_margin", "stock_market", "stock_reference",
		"industry_tmt", "macro_business", "macro_economy",
		"macro_interest_rate", "macro_price", "wealth_fund_sales",
	}
	for i, module := range modules {
		fmt.Printf("  - %s\n", module)
		if (i+1)%5 == 0 && i != len(modules)-1 {
			fmt.Println()
		}
	}

	fmt.Println("\n💡 使用提示:")
	fmt.Println("  - 服务器支持 CORS，可从浏览器直接调用")
	fmt.Println("  - 所有 API 支持 JSON 请求和响应")
	fmt.Println("  - 统一端点提供更简洁的 API 结构")
	fmt.Println("  - 详细文档请查看 docs/HTTP_UNIFIED_ENDPOINTS.md")

	fmt.Println("\n🛑 按 Ctrl+C 停止服务器")

	// 等待用户中断
	<-ctx.Done()
	fmt.Println("\n👋 服务器已停止")
}