package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func BenchmarkInefficientCompute(b *testing.B) {
	router := gin.Default()
	router.GET("/compute", inefficientCompute)

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/compute", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != 200 {
			b.Errorf("Expected status 200, got %v", resp.Code)
		}
	}
}
