.PHONY: test build clean gen examples

test:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

build-mcp:
	go build -o bin/tushare-mcp ./cmd/mcp-server

build-gen:
	go build -o bin/generator ./cmd/generator

gen:
	./bin/generator pkg/sdk/api

examples: build-examples

build-examples:
	go build -o bin/stock_basic_example ./cmd/examples/stock_basic
	go build -o bin/daily_example ./cmd/examples/daily
	go build -o bin/daily_basic_example ./cmd/examples/daily_basic
	go build -o bin/trade_cal_example ./cmd/examples/trade_cal

clean:
	rm -rf bin/
	rm -f coverage.out coverage.html
