package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRegisterFrontendRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	distPath := t.TempDir()
	if err := os.WriteFile(filepath.Join(distPath, "index.html"), []byte("frontend index"), 0644); err != nil {
		t.Fatal(err)
	}

	assetsPath := filepath.Join(distPath, "assets")
	if err := os.MkdirAll(assetsPath, 0755); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(filepath.Join(assetsPath, "app.js"), []byte("frontend asset"), 0644); err != nil {
		t.Fatal(err)
	}

	router := gin.New()
	registerFrontendRoutes(router, distPath)

	tests := []struct {
		name           string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "static file",
			path:           "/assets/app.js",
			expectedStatus: http.StatusOK,
			expectedBody:   "frontend asset",
		},
		{
			name:           "spa route",
			path:           "/diary/entry/123",
			expectedStatus: http.StatusOK,
			expectedBody:   "frontend index",
		},
		{
			name:           "api route",
			path:           "/api/missing",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "NOT_FOUND",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, test.path, nil)
			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, request)

			if recorder.Code != test.expectedStatus {
				t.Fatalf("expected status %d, got %d", test.expectedStatus, recorder.Code)
			}

			if !strings.Contains(recorder.Body.String(), test.expectedBody) {
				t.Fatalf("expected response body to contain %q, got %q", test.expectedBody, recorder.Body.String())
			}
		})
	}
}

func TestRegisterFrontendRoutesDoesNotUseLaterAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	distPath := t.TempDir()
	if err := os.WriteFile(filepath.Join(distPath, "index.html"), []byte("frontend index"), 0644); err != nil {
		t.Fatal(err)
	}

	router := gin.New()
	registerFrontendRoutes(router, distPath)

	authenticated := router.Group("", func(ctx *gin.Context) {
		ctx.Status(http.StatusUnauthorized)
		ctx.Abort()
	})
	authenticated.GET("/api/protected", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	frontendRequest := httptest.NewRequest(http.MethodGet, "/", nil)
	frontendRecorder := httptest.NewRecorder()
	router.ServeHTTP(frontendRecorder, frontendRequest)

	if frontendRecorder.Code != http.StatusOK {
		t.Fatalf("expected frontend status %d, got %d", http.StatusOK, frontendRecorder.Code)
	}

	if !strings.Contains(frontendRecorder.Body.String(), "frontend index") {
		t.Fatalf("expected frontend response body, got %q", frontendRecorder.Body.String())
	}

	apiRequest := httptest.NewRequest(http.MethodGet, "/api/protected", nil)
	apiRecorder := httptest.NewRecorder()
	router.ServeHTTP(apiRecorder, apiRequest)

	if apiRecorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected protected API status %d, got %d", http.StatusUnauthorized, apiRecorder.Code)
	}
}
