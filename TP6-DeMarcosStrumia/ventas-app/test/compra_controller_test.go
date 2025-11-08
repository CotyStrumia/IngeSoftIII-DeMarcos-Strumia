package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"ventas-app/controllers"
	"ventas-app/database"
	"ventas-app/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegistrarCompra_InvalidBody(t *testing.T) {
	database.GetDB = func(c *gin.Context) database.DBHandler {
		return &mocks.MockDB{ShouldErr: false}
	}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/compras", controllers.RegistrarCompra)

	body := `{"producto_id": 0, "cantidad": 0}`
	req, _ := http.NewRequest("POST", "/compras", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
