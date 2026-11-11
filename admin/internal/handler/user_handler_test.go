package handler

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"linkspan/internal/common/convention/result"
	"linkspan/internal/middleware"

	"github.com/gin-gonic/gin"
)

func TestUserHandlerRejectsInvalidInputBeforeService(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewUserHandler(nil) // 无效输入必须在访问业务依赖前被拒绝。
	r := gin.New()
	r.Use(middleware.ErrorHandler())
	r.POST("/login", h.Login)
	r.POST("/register", h.Register)
	r.PUT("/user", h.Update)
	r.GET("/check", h.CheckLogin)
	r.GET("/has", h.HasUsername)
	r.DELETE("/logout", h.Logout)
	for _, tt := range []struct{ name, method, path, body, code, message string }{
		{"empty_body", "POST", "/login", "", "A000004", "请求体不能为空"},
		{"invalid_json", "POST", "/login", `{"username":`, "A000002", "请求体格式错误"},
		{"missing_password", "POST", "/login", `{"username":"test3"}`, "A000003", "密码不能为空"},
		{"wrong_type", "POST", "/login", `{"username":123}`, "A000002", "用户名类型错误"},
		{"register_validation", "POST", "/register", `{}`, "A000003", "用户名不能为空"},
		{"update_null", "PUT", "/user", `null`, "A000003", "用户名不能为空"},
		{"update_empty", "PUT", "/user", `{}`, "A000003", "用户名不能为空"},
		{"check_query", "GET", "/check?username=test3", "", "A000003", "登录凭证不能为空"},
		{"has_query", "GET", "/has", "", "A000003", "用户名不能为空"},
		{"logout_query", "DELETE", "/logout?username=test3", "", "A000003", "登录凭证不能为空"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, tt.path, strings.NewReader(tt.body))
			request.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, request)
			var response result.Result[any]
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}
			if w.Code != 200 || response.Code != tt.code || response.Message != tt.message {
				t.Fatalf("unexpected response: status=%d body=%s", w.Code, w.Body.String())
			}
		})
	}
}
