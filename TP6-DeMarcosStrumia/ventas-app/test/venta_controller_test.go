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

func TestRegistrarVenta_InvalidBody(t *testing.T) {
	database.GetDB = func(c *gin.Context) database.DBHandler {
		return &mocks.MockDB{ShouldErr: false}
	}

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/ventas", controllers.RegistrarVenta)

	// Arrange: sin datos de producto ni cantidad
	body := `{"producto_id": 0, "cantidad": 0}`
	req, _ := http.NewRequest("POST", "/ventas", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	// Act
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.Code, "Debe devolver 400 si los datos son inválidos")
}
