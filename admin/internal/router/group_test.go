package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"linkspan/admin/internal/common/biz/usercontext"
	"linkspan/admin/internal/handler"
	"linkspan/admin/internal/middleware"
	"linkspan/admin/internal/service"

	"github.com/gin-gonic/gin"
)

func TestGroupRouteRejectsRequestsBeforeStorage(t *testing.T) {
	tests := []struct {
		name, body, code string
		authenticated    bool
	}{
		{"missing_login", `{"name":"学习资料"}`, "A000400", false},
		{"malformed_json", `{"name":`, "A000002", true},
		{"missing_name", `{}`, "A000003", true},
		{"null_body", `null`, "A000003", true},
		{"wrong_type", `{"name":123}`, "A000002", true},
		{"too_long", `{"name":"` + strings.Repeat("中", 65) + `"}`, "A000003", true},
		{"whitespace", `{"name":"   "}`, "A000500", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.Use(middleware.ErrorHandler())
			auth := middleware.RequireLogin(nil)
			if tt.authenticated {
				// 参数校验测试注入可信上下文，不连接实际 Redis。
				auth = func(c *gin.Context) {
					c.Request = c.Request.WithContext(usercontext.WithUsername(c.Request.Context(), "test3"))
					c.Next()
				}
			}
			Register(r, &handler.UserHandler{}, handler.NewGroupHandler(&service.GroupService{}), auth)
			request := httptest.NewRequest(http.MethodPost, "/api/short-link/admin/v1/group", strings.NewReader(tt.body))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			r.ServeHTTP(response, request)
			var result struct {
				Code      string `json:"code"`
				Success   bool   `json:"success"`
				RequestID string `json:"requestId"`
			}
			if err := json.Unmarshal(response.Body.Bytes(), &result); err != nil {
				t.Fatal(err)
			}
			if response.Code != http.StatusOK || result.Code != tt.code || result.Success || result.RequestID == "" {
				t.Fatalf("unexpected response: status=%d body=%s", response.Code, response.Body.String())
			}
		})
	}
}
