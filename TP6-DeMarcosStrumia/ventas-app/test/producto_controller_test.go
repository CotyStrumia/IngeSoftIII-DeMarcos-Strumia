package test

import (
	_ "bytes"
	_ "encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"ventas-app/controllers"
	"ventas-app/database"
	"ventas-app/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestListarProductos_OK(t *testing.T) {
	database.GetDB = func(c *gin.Context) database.DBHandler {
		return &mocks.MockDB{ShouldErr: false}
	}

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/productos", controllers.ListarProductos)

	req, _ := http.NewRequest("GET", "/productos", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
