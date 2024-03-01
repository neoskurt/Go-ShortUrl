// Dans un fichier user_test.go
package controllers

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
	"go-url-shortener/controllers"
	"go-url-shortener/models"

    
)

func TestRegisterUser(t *testing.T) {
    models.InitDB()
    gin.SetMode(gin.TestMode)
    r := gin.Default()

	r.POST("/register", controllers.RegisterUser)

    var jsonStr = []byte(`{
		"email": "test@example.com",
		"password": "password",
		"username": "exemple"
	}`)
    req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonStr))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
    }

}
