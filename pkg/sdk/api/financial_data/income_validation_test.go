package api

import (
	"testing"
)

func TestIncomeRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     *IncomeRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "有效请求 - ts_code",
			req: &IncomeRequest{
				TsCode: "000001.SZ",
			},
			wantErr: false,
		},
		{
			name: "有效请求 - report_type",
			req: &IncomeRequest{
				TsCode:     "000001.SZ",
				ReportType: "0",
			},
			wantErr: false,
		},
		{
			name: "无效请求 - ts_code太短",
			req: &IncomeRequest{
				TsCode: "000001",
			},
			wantErr: true,
			errMsg:  "ts_code 必须为9个字符",
		},
		{
			name: "无效请求 - report_type错误",
			req: &IncomeRequest{
				TsCode:     "000001.SZ",
				ReportType: "9",
			},
			wantErr: true,
			errMsg:  "report_type 必须为 0、1、2、3、4、5 之一",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if tt.wantErr {
				if err == nil {
					t.Errorf("期望错误，但得到nil")
					return
				}
				if tt.errMsg != "" && err.Error()[:len(tt.errMsg)] != tt.errMsg {
					t.Errorf("错误消息应包含%q，得到%q", tt.errMsg, err.Error())
				}
				t.Logf("✓ 正确拒绝无效请求: %v", err)
			} else {
				if err != nil {
					t.Errorf("意外错误: %v", err)
				} else {
					t.Log("✓ 有效请求被接受")
				}
			}
		})
	}
}