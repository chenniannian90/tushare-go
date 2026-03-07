package api

import (
	"testing"
)

func TestIndexBasicRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     *IndexBasicRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "有效请求 - ts_code",
			req: &IndexBasicRequest{
				TsCode: "000001.SH",
			},
			wantErr: false,
		},
		{
			name: "有效请求 - market",
			req: &IndexBasicRequest{
				Market: "SSE",
			},
			wantErr: false,
		},
		{
			name: "无效请求 - ts_code太短",
			req: &IndexBasicRequest{
				TsCode: "000001",
			},
			wantErr: true,
			errMsg:  "ts_code 必须为9个字符",
		},
		{
			name: "无效请求 - ts_code后缀错误",
			req: &IndexBasicRequest{
				TsCode: "000001.XY",
			},
			wantErr: true,
			errMsg:  "ts_code 必须以.SZ或.SH结尾",
		},
		{
			name: "无效请求 - market错误",
			req: &IndexBasicRequest{
				Market: "INVALID",
			},
			wantErr: true,
			errMsg:  "market 必须为 SSE、SZSE、CZCE、SHFE、INE、BJE 之一",
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