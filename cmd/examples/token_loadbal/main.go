// +build ignore

package main

import (
	"fmt"
	"log"

	"tushare-go/pkg/sdk"
)

func main() {
	fmt.Println("=== Token 负载均衡示例 ===\n")

	// 示例1: 使用多个 token，自动轮询负载均衡
	fmt.Println("1. 多 Token 负载均衡示例:")
	tokens := []string{
		"your-token-1",
		"your-token-2",
		"your-token-3",
	}

	config, err := sdk.NewConfigWithTokens(tokens)
	if err != nil {
		log.Fatalf("创建配置失败: %v", err)
	}

	_ = sdk.NewClient(config) // 创建客户端（演示用）

	// 模拟多次请求，观察 token 轮询
	fmt.Println("连续10次请求，观察 token 轮询选择:")
	for i := 1; i <= 10; i++ {
		// 获取当前使用的token
		currentToken := config.GetToken()
		fmt.Printf("  请求 %d: 使用 token %s\n", i, getShortToken(currentToken))
	}

	fmt.Println("\n2. 向后兼容示例（单个 token）:")
	config2, err := sdk.NewConfig("your-single-token")
	if err != nil {
		log.Fatalf("创建配置失败: %v", err)
	}

	fmt.Println("连续3次请求，观察单个 token 使用:")
	for i := 1; i <= 3; i++ {
		currentToken := config2.GetToken()
		fmt.Printf("  请求 %d: 使用 token %s\n", i, getShortToken(currentToken))
	}

	fmt.Println("\n=== 负载均衡说明 ===")
	fmt.Println("多 Token (自动轮询): 按顺序依次使用每个 token，实现负载均衡")
	fmt.Println("单个 Token: 传统的使用方式，向后兼容")
	fmt.Println("")
	fmt.Println("优势:")
	fmt.Println("- 充分利用多个 API 配额")
	fmt.Println("- 避免单个 token 限流")
	fmt.Println("- 提高整体吞吐量")
	fmt.Println("- 自动负载分配，无需手动切换")
	fmt.Println("- 并发安全，使用原子操作保证")
}

// getShortToken 返回 token 的简短形式（用于演示）
func getShortToken(token string) string {
	if len(token) <= 10 {
		return token
	}
	return token[:10] + "..."
}
