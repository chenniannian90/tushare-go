// This example demonstrates how to use the logging system in tushare-go
package main

import (
	"context"
	"io"
	"os"
	"time"

	"tushare-go/pkg/sdk"
	"tushare-go/pkg/sdk/logger"
	futures "tushare-go/pkg/sdk/api/futures"
)

func main() {
	// 示例1：基础配置 - 只输出到控制台
	basicExample()

	// 示例2：文件配置 - 输出到文件
	// fileExample()

	// 示例3：高级配�� - 同时输出到控制台和文件
	advancedExample()
}

// 基础配置示例
func basicExample() {
	println("\n=== 示例1：基础配置（控制台输出）===")

	logger.Init(&logger.LogConfig{
		Filename: "",     // 留空表示输出到控制台
		Level:    "debug", // 显示所有级别
		Format:   "text",
	})

	logger.Info("日志系统已启动（控制台模式）")
	logger.Debug("这是调试信息")
	logger.Warn("这是警告信息")

	config, _ := sdk.NewConfig("your-token")
	client := sdk.NewClient(config)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := futures.TradeCal(ctx, client, &futures.TradeCalRequest{
		Exchange: "SHFE",
		StartDate: "20250101",
		EndDate:   "20250103",
	})

	if err != nil {
		logger.WithError(err).Error("API调用失败")
	} else {
		logger.Info("API调用成功")
	}
}

// 文件配置示例
func fileExample() {
	println("\n=== 示例2：文件配置（输出到文件）===")

	logger.Init(&logger.LogConfig{
		Filename:   "app.log",      // 输出到文件
		MaxSize:    10,             // 10MB
		MaxAge:     30,             // 保留30天
		MaxBackups: 3,              // 保留3个备份
		Compress:   true,           // 压缩旧文件
		Level:      "info",
		Format:     "text",
	})

	logger.Info("日志系统已启动（文件模式）")
	logger.Info("日志文件：app.log")

	// 运行一些操作...
	logger.Info("执行业务逻辑...")
	time.Sleep(100 * time.Millisecond)
	logger.Info("业务逻辑完成")

	// 查看日志文件
	println("请查看 app.log 文件查看日志内容")
}

// 高级配置示例
func advancedExample() {
	println("\n=== 示例3：高级配置（同时输出到控制台和文件）===")

	// 创建日志文件
	logFile, err := os.OpenFile("advanced.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	// 同时输出到控制台和文件
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logger.SetOutput(multiWriter)

	// 初始化日志系统
	logger.Init(&logger.LogConfig{
		Level:  "debug",
		Format: "text",
	})

	logger.Info("========================================")
	logger.Info("高级日志配置示例")
	logger.Info("========================================")

	logger.Info("日志同时输出到：")
	logger.Info("  1. 控制台（屏幕）")
	logger.Info("  2. 文件（advanced.log）")

	// 演示结构化日志
	logger.WithFields(logger.Fields{
		"service": "tushare-api",
		"version": "1.0.0",
	}).Info("服务启动")

	// 模拟业务流程
	processBusinessLogic()

	logger.Info("========================================")
	logger.Info("示例完成")
	logger.Info("========================================")

	// 提示查看日志文件
	println("\n日志已保存到：advanced.log")
	println("同时显示在控制台")
}

// 模拟业务逻辑
func processBusinessLogic() {
	logger.Info("开始处理业务逻辑")

	// 步骤1
	logger.WithField("step", 1).Info("初始化配置")
	time.Sleep(50 * time.Millisecond)

	// 步骤2
	logger.WithField("step", 2).Info("连接 API 服务器")
	time.Sleep(50 * time.Millisecond)

	// 步骤3
	logger.WithFields(logger.Fields{
		"step":     3,
		"api_name": "trade_cal",
		"exchange": "SHFE",
	}).Info("获取交易日历")

	token := "412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1"
	config, _ := sdk.NewConfig(token)
	client := sdk.NewClient(config)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	calendar, err := futures.TradeCal(ctx, client, &futures.TradeCalRequest{
		Exchange: "SHFE",
		StartDate: "20250101",
		EndDate:   "20250103",
	})

	if err != nil {
		logger.WithError(err).Error("获取交易日历失败")
		return
	}

	logger.WithField("count", len(calendar)).Info("获取交易日历成功")

	// 显示部分数据
	for i, item := range calendar {
		if i >= 2 {
			break
		}
		status := "休市"
		if item.IsOpen == "1" {
			status = "交易"
		}
		logger.WithFields(logger.Fields{
			"date":     item.CalDate,
			"status":   status,
			"exchange": item.Exchange,
		}).Info("交易日历数据")
	}

	// 完成
	logger.WithField("step", "final").Info("业务逻辑处理完成")
}
