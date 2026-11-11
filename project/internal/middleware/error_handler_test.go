package middleware_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"linkspan/internal/middleware"
	"linkspan/project/internal/handler"
	"linkspan/project/internal/router"
	"linkspan/project/internal/service"

	"github.com/gin-gonic/gin"
)

func TestShortLinkErrorResponses(t *testing.T) {
	engine := gin.New()
	engine.Use(middleware.ErrorHandler())
	router.Register(engine, handler.NewShortLinkHandler(service.NewShortLinkService(nil, nil, nil)))
	engine.GET("/timeout", func(c *gin.Context) { _ = c.Error(context.DeadlineExceeded) })
	engine.GET("/failure", func(c *gin.Context) { _ = c.Error(errors.New("internal detail")) })
	engine.GET("/panic", func(c *gin.Context) { panic("internal detail") })

	for _, tt := range []struct {
		method, path, body, code, message string
	}{
		{"POST", "/api/short-link/v1/create", "{", "A000002", "请求体格式错误"},
		{"GET", "/api/short-link/v1/page", "", "A000003", "分组标识不能为空"},
		{"POST", "/api/short-link/v1/update", "{", "A000002", "请求体格式错误"},
		{"POST", "/api/short-link/v1/update", `{"gid":"group","originGid":"group","fullShortUrl":"localhost:8002/12345678","originUrl":"invalid","validDateType":0}`, "A000003", "originUrl 须为完整的 HTTP 或 HTTPS 地址"},
		{"GET", "/timeout", "", "B000100", "系统执行超时"},
		{"GET", "/failure", "", "B000001", "系统执行错误"},
		{"GET", "/panic", "", "B000001", "系统执行错误"},
	} {
		t.Run(tt.path+"/"+tt.message, func(t *testing.T) {
			w := httptest.NewRecorder()
			engine.ServeHTTP(w, httptest.NewRequest(tt.method, tt.path, strings.NewReader(tt.body)))

			var reply struct {
				Code      string
				Message   string
				RequestID string
				Success   bool
				Data      json.RawMessage
			}
			if err := json.Unmarshal(w.Body.Bytes(), &reply); err != nil {
				t.Fatalf("invalid response: %v, body=%s", err, w.Body.String())
			}
			if w.Code != http.StatusOK || reply.Code != tt.code || reply.Message != tt.message ||
				reply.Success || string(reply.Data) != "null" || reply.RequestID == "" ||
				reply.RequestID != w.Header().Get("X-Request-ID") {
				t.Fatalf("unexpected response: status=%d, body=%s", w.Code, w.Body.String())
			}
		})
	}
}
