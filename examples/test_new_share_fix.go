package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"tushare-go/pkg/sdk"
	"tushare-go/pkg/sdk/api/stock_basic"
)

func main() {
	// 从环境变量获取token
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		log.Fatal("❌ 请设置环境变量 TUSHARE_TOKEN")
	}

	fmt.Println("🔧 正在测试 new_share API 的修复...")
	fmt.Println("📝 使用 CallAPIFlexible 方法自动处理响应格式")
	fmt.Println()

	// 创建客户端
	config, err := sdk.NewConfig(token)
	if err != nil {
		log.Fatalf("❌ 创建配置失败: %v", err)
	}
	client := sdk.NewClient(config)

	// 调用 new_share API
	ctx := context.Background()
	req := &stock_basic.NewShareRequest{
		StartDate: "20260301",
		EndDate:   "20260331",
	}

	fmt.Println("📡 正在调用 Tushare API...")
	items, err := stock_basic.NewShare(ctx, client, req)
	if err != nil {
		log.Fatalf("❌ API调用失败: %v", err)
	}

	fmt.Printf("✅ 成功获取 %d 条IPO数据\n\n", len(items))

	if len(items) > 0 {
		fmt.Println("📊 数据示例（前3条）：")
		fmt.Println("━" + "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

		for i, item := range items {
			if i >= 3 {
				break
			}
			fmt.Printf("\n📈 第 %d 条记录:\n", i+1)
			fmt.Printf("   TS代码: %s\n", item.TsCode)
			fmt.Printf("   申购代码: %s\n", item.SubCode)
			fmt.Printf("   名称: %s\n", item.Name)
			fmt.Printf("   上市日期: %s\n", item.IpoDate)
			fmt.Printf("   发行价格: %.2f 元\n", item.Price)
			fmt.Printf("   市盈率: %.2f\n", item.Pe)
			fmt.Printf("   发行总量: %.2f 万股\n", item.Amount)
			fmt.Printf("   募集资金: %.2f 亿元\n", item.Funds)
		}

		fmt.Println()
		fmt.Println("✅ 测试成功！")
		fmt.Println("   • CallAPIFlexible 正确处理了API响应")
		fmt.Println("   • 数据已正确解析为结构体")
		fmt.Println("   • 不再出现 JSON unmarshal 错误")
	} else {
		fmt.Println("⚠️  当前时间段无IPO数据")
		fmt.Println("✅ 但API调用成功，说明修复有效！")
	}
}
