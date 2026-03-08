// +build ignore

package main

import (
	"fmt"

	"tushare-go/pkg/mcp/server"
)

func main() {
	fmt.Println("=== HTTP 统一端点示例 ===\n")

	// 创建 HTTP 路由器
	router := server.NewHTTPRouter()

	// 1. 展示传统路径方式
	fmt.Println("📋 传统路径方式（向后兼容）:")
	fmt.Println("  POST /api/v1/hk_stock/hk_basic")
	fmt.Println("  POST /api/v1/hk_stock/hk_daily")
	fmt.Println("  POST /api/v1/hk_stock/hk_cal")
	fmt.Println("  POST /api/v1/hk_stock/hk_min")
	fmt.Println("  POST /api/v1/hk_stock/hk_factor")
	fmt.Println()

	// 2. 展示统一端点方式
	fmt.Println("🎯 统一端点方式（推荐）:")
	fmt.Println("  POST /api/v1/hk_stock?tool=hk_basic")
	fmt.Println("  POST /api/v1/hk_stock?tool=hk_daily")
	fmt.Println("  POST /api/v1/hk_stock?tool=hk_cal")
	fmt.Println("  POST /api/v1/hk_stock?tool=hk_min")
	fmt.Println("  POST /api/v1/hk_stock?tool=hk_factor")
	fmt.Println()

	// 3. 查询统一端点信息
	fmt.Println("🔍 统一端点详情:")
	unifiedEndpoints := router.GetUnifiedEndpoints()
	for _, endpoint := range unifiedEndpoints {
		fmt.Printf("  路径: %s\n", endpoint.Path)
		fmt.Printf("  模块: %s\n", endpoint.Module)
		fmt.Printf("  参数名: %s\n", endpoint.ParamName)
		fmt.Printf("  描述: %s\n", endpoint.Description)
	}
	fmt.Println()

	// 4. 展示 hk_stock 模块的所有工具
	fmt.Println("🛠️  HK Stock 模块可用工具:")
	hkTools := router.GetModuleTools("hk_stock")
	for i, tool := range hkTools {
		fmt.Printf("  %d. %s\n", i+1, tool)
	}
	fmt.Println()

	// 5. 演示如何解析统一端点调用
	fmt.Println("🔧 统一端点解析示例:")

	testCases := []struct {
		path   string
		tool   string
		desc   string
	}{
		{"hk_stock", "hk_basic", "获取基础信息"},
		{"hk_stock", "hk_daily", "获取日线数据"},
		{"hk_stock", "hk_cal", "获取交易日历"},
		{"hk_stock", "invalid_tool", "无效工具（错误示例）"},
	}

	for _, tc := range testCases {
		fmt.Printf("  调用: %s?tool=%s (%s)\n", tc.path, tc.tool, tc.desc)

		toolName, err := router.ResolveUnifiedEndpoint(tc.path, tc.tool)
		if err != nil {
			fmt.Printf("    ❌ 错误: %v\n", err)
		} else {
			fmt.Printf("    ✅ 解析为: %s\n", toolName)
		}
	}
	fmt.Println()

	// 6. 展示路由统计
	fmt.Println("📊 路由统计信息:")
	hkRoutes := router.GetModuleRoutes("hk_stock")
	fmt.Printf("  HK Stock 模块总路由数: %d\n", len(hkRoutes))
	fmt.Printf("    - 单独工具路由: %d\n", len(hkTools))
	fmt.Printf("    - 统一端点路由: 1\n")
	fmt.Println()

	// 7. 对比不同模块
	fmt.Println("🔄 模块对比:")

	modules := []string{"hk_stock", "bond", "etf"}
	for _, module := range modules {
		tools := router.GetModuleTools(module)
		_, hasUnified := router.GetModuleUnifiedEndpoint(module)

		status := "❌"
		if hasUnified {
			status = "✅"
		}

		fmt.Printf("  %s %s: %d 个工具\n", status, module, len(tools))
	}
	fmt.Println()

	// 8. API 调用示例
	fmt.Println("💡 cURL 调用示例:")
	fmt.Println("  # 使用统一端点")
	fmt.Println(`  curl -X POST "http://localhost:8080/api/v1/hk_stock?tool=hk_basic" \`)
	fmt.Println(`    -H "Content-Type: application/json" \`)
	fmt.Println(`    -d '{"ts_code": "00700.HK", "list_date": "20240101"}'`)
	fmt.Println()

	fmt.Println("  # 使用传统路径")
	fmt.Println(`  curl -X POST "http://localhost:8080/api/v1/hk_stock/hk_basic" \`)
	fmt.Println(`    -H "Content-Type: application/json" \`)
	fmt.Println(`    -d '{"ts_code": "00700.HK", "list_date": "20240101"}'`)
	fmt.Println()

	// 9. 错误处理示例
	fmt.Println("⚠️  错误处理示例:")
	fmt.Println("  无效工具: POST /api/v1/hk_stock?tool=invalid_tool")
	fmt.Println("  响应: {")
	fmt.Println(`    "error": {`)
	fmt.Println(`      "code": -32601,`)
	fmt.Println(`      "message": "tool 'invalid_tool' not found in module 'hk_stock'. Available tools: [...]"`)
	fmt.Println(`    }`)
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("✅ 演示完成！")
}
