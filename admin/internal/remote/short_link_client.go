package remote

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"linkspan/admin/internal/common/convention/result"
	"net/http"
	"time"
)

type ShortLinkClient struct {
	client *http.Client
}

func NewShortLinkClient() *ShortLinkClient {
	return &ShortLinkClient{client: &http.Client{Timeout: 8 * time.Second}} 
}

func (r *ShortLinkClient) Call(ctx context.Context, method, path string, payload any) (result.Result[json.RawMessage], error) {
	var reply result.Result[json.RawMessage]
	var body io.Reader

	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return reply, fmt.Errorf("编码核心请求失败: %w", err)
		}
		body = bytes.NewReader(data)
	}

	request, err := http.NewRequestWithContext(ctx, method, "http://127.0.0.1:8002/api/short-link/v1"+path, body)
	if err != nil {
		return reply, fmt.Errorf("构造核心请求失败: %w", err)
	}

	if payload != nil {
		request.Header.Set("Content-Type", "application/json")
	}	

	response, err := r.client.Do(request)
	if err != nil {
		return reply, fmt.Errorf("调用短链接核心服务失败: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return reply, fmt.Errorf("核心服务返回 HTTP %d", response.StatusCode)
	}
	if err := json.NewDecoder(io.LimitReader(response.Body, 2<<20)).Decode(&reply); err != nil {
		return reply, fmt.Errorf("解析核心响应失败: %w", err)
	}
	if reply.Code == "" || reply.Success != (reply.Code == result.SuccessCode) {
		return reply, fmt.Errorf("核心服务响应格式无效")
	}

	return reply, nil	
}