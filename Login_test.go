// Dans un fichier user_test.go
package controllers

import (
    "io"
    "bytes"
    "net/http"
    "net/http/httptest"
    "encoding/json"
    "testing"
    "github.com/gin-gonic/gin"
	"go-url-shortener/controllers"
    "go-url-shortener/routes"
	"go-url-shortener/models"
)

func TestRegisterUser(t *testing.T) {
    models.InitDB()
    gin.SetMode(gin.TestMode) 

    r := routes.InitializeRoutes(gin.Default())

    user := models.LoginCredentials{
        Email: "test@example.com",
        Password: "password",
        Username: "example",
    }

    jsonLogin , err := json.Marshal(user)
    if err != nil {
        t.Fatal(err)
    }
    reqRes, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonLogin))
    if err != nil {
        t.Fatal(err)
    }
    reqRes.Header.Set("Content-Type", "application/json")

    wR := httptest.NewRecorder()
    r.ServeHTTP(wR, reqRes)

    if wR.Code != http.StatusOK {
        t.Errorf("Expected status code %d, but got %d", http.StatusOK, wR.Code)
    }

 
    reqLog, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonLogin))
    if err != nil {
        t.Fatal(err)
    }
    reqLog.Header.Set("Content-Type", "application/json")

    wL := httptest.NewRecorder()
    r.ServeHTTP(wL, reqLog)
    if wL.Code != http.StatusOK {
        body, _ := io.ReadAll(wL.Body)

        t.Errorf("Expected status code %d, got %d, %s", http.StatusOK, wL.Code, string(body))
    }

}


