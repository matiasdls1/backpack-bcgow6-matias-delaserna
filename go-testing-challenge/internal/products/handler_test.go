package products

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	repo := NewRepository()
	svc := NewService(repo)
	handler := NewHandler(svc)
	router := gin.Default()

	router.GET("/api/v1/products", handler.GetProducts)
	return router
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestGetProducts(t *testing.T) {
	server := createServer()
	req, resprec := createRequestTest(http.MethodGet, "/api/v1/products?seller_id=foo", "")
	server.ServeHTTP(resprec, req)
	assert.Equal(t, resprec.Code, 200)
}

func TestGetProductsBad(t *testing.T) {
	server := createServer()
	req, resprec := createRequestTest(http.MethodGet, "/api/v1/products", "")
	server.ServeHTTP(resprec, req)
	assert.Equal(t, resprec.Code, 400)
}
