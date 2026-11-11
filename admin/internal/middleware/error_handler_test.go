package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http/httptest"
	"strings"
	"testing"

	"linkspan/admin/internal/common/convention/errorcode"
	"linkspan/internal/common/convention/exception"
	"linkspan/internal/common/convention/result"
	httpmiddleware "linkspan/internal/middleware"

	"github.com/gin-gonic/gin"
)

func TestErrorHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var logs bytes.Buffer
	previous := log.Writer()
	log.SetOutput(&logs)
	t.Cleanup(func() { log.SetOutput(previous) })
	engine := gin.New()
	engine.Use(httpmiddleware.RequestLogger(), httpmiddleware.ErrorHandler())
	cause := errors.New("diagnostic-db-cause")
	engine.GET("/business", func(c *gin.Context) {
		_ = c.Error(fmt.Errorf("service: %w", exception.WrapServiceError(errorcode.UserNull, cause)))
	})
	engine.GET("/unknown", func(c *gin.Context) { _ = c.Error(cause) })
	engine.GET("/timeout", func(c *gin.Context) { _ = c.Error(fmt.Errorf("database: %w", context.DeadlineExceeded)) })
	engine.GET("/panic", func(c *gin.Context) { panic("private-panic-value") })
	engine.GET("/written", func(c *gin.Context) {
		c.JSON(201, gin.H{"already": "written"})
		_ = c.Error(cause)
	})
	engine.GET("/healthy", func(c *gin.Context) { c.JSON(200, gin.H{"ok": true}) })
	engine.GET("/protected", RequireLogin(nil), func(c *gin.Context) { t.Error("unauthenticated request reached handler") })
	for _, tt := range []struct{ path, code string }{
		{"/business", "B000200"}, {"/unknown", "B000001"},
		{"/timeout", "B000100"}, {"/panic", "B000001"}, {"/protected", "A000400"},
	} {
		t.Run(tt.path, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", tt.path+"?token=private-query-token", nil)
			r.Header.Set("token", "private-header-token")
			engine.ServeHTTP(w, r)
			var response result.Result[any]
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}
			if w.Code != 200 || response.Code != tt.code || response.Success {
				t.Fatalf("unexpected error response: status=%d body=%s", w.Code, w.Body.String())
			}
			if response.RequestID == nil || *response.RequestID == "" || *response.RequestID != w.Header().Get("X-Request-ID") {
				t.Fatal("request ID missing or inconsistent")
			}
			if strings.Contains(w.Body.String(), cause.Error()) {
				t.Fatal("internal cause leaked into response")
			}
		})
	}
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, httptest.NewRequest("GET", "/written", nil))
	if w.Code != 201 || w.Body.String() != `{"already":"written"}` {
		t.Fatalf("response written twice: %s", w.Body.String())
	}
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, httptest.NewRequest("GET", "/healthy", nil))
	if w.Code != 200 || w.Body.String() != `{"ok":true}` {
		t.Fatal("server did not recover after panic")
	}
	if !strings.Contains(logs.String(), cause.Error()) {
		t.Fatal("original cause missing from logs")
	}
	for _, secret := range []string{"private-query-token", "private-header-token", "private-panic-value"} {
		if strings.Contains(logs.String(), secret) {
			t.Fatalf("logs leaked %s", secret)
		}
	}
}
