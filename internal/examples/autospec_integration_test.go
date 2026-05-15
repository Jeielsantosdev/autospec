package examples_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jeielsantosdev/autospec/internal/autospec"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
	"github.com/gin-gonic/gin"
)

func TestAutospecGeneratesOpenAPISpecForBasicRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()

	// Attach autospec (autodetecta Gin e instala middleware + /openapi.json)
	as := autospec.New()
	if err := as.AttachTo(engine); err != nil {
		t.Fatalf("failed to attach autospec: %v", err)
	}

	// Register minimal handlers
	engine.POST("/users", func(c *gin.Context) {
		var b map[string]any
		_ = c.BindJSON(&b)
		c.JSON(http.StatusCreated, gin.H{"id": 1})
	})
	engine.PUT("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"updated": true})
	})
	engine.DELETE("/users/:id", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
	engine.GET("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"id": c.Param("id")})
	})

	// Fire requests to generate runtime samples (capture middleware depends on actual requests)
	tests := []struct {
		method string
		url    string
		body   string
		code   int
	}{
		{"POST", "/users", `{"name":"john"}`, http.StatusCreated},
		{"PUT", "/users/1", `{"name":"john updated"}`, http.StatusOK},
		{"DELETE", "/users/1", ``, http.StatusNoContent},
		{"GET", "/users/1", ``, http.StatusOK},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(tt.method, tt.url, bytes.NewBufferString(tt.body))
		if tt.body != "" {
			req.Header.Set("Content-Type", "application/json")
		}
		w := httptest.NewRecorder()
		engine.ServeHTTP(w, req)
		if w.Code != tt.code {
			t.Fatalf("expected status %d for %s %s, got %d", tt.code, tt.method, tt.url, w.Code)
		}
	}

	// Fetch generated openapi.json
	req := httptest.NewRequest("GET", "/openapi.json", nil)
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("/openapi.json returned status %d", w.Code)
	}

	var spec openapi.Spec
	if err := json.Unmarshal(w.Body.Bytes(), &spec); err != nil {
		t.Fatalf("failed to parse openapi.json: %v; body: %s", err, string(w.Body.Bytes()))
	}

	// Checks: POST /users exists
	p1, ok := spec.Paths["/users"]
	if !ok {
		t.Fatalf("spec missing path: /users")
	}
	if p1.Post == nil {
		t.Fatalf("spec missing POST operation for /users")
	}

	// Checks: GET/PUT/DELETE for /users/:id exist
	p2, ok := spec.Paths["/users/:id"]
	if !ok {
		t.Fatalf("spec missing path: /users/:id")
	}
	if p2.Get == nil {
		t.Fatalf("spec missing GET operation for /users/:id")
	}
	if p2.Put == nil {
		t.Fatalf("spec missing PUT operation for /users/:id")
	}
	if p2.Delete == nil {
		t.Fatalf("spec missing DELETE operation for /users/:id")
	}
}
