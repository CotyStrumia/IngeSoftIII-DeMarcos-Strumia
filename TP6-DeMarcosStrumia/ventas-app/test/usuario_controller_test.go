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

func TestCrearUsuario_InvalidData(t *testing.T) {

	database.GetDB = func(c *gin.Context) database.DBHandler {
		return &mocks.MockDB{ShouldErr: false}
	}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/usuarios", controllers.CrearUsuario)

	body := `{"nombre": "", "clave": "", "rol": ""}`
	req, _ := http.NewRequest("POST", "/usuarios", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
