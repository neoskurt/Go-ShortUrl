// Dans un fichier user_test.go
package controllers

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
	"go-url-shortener/models"
    "go-url-shortener/routes"

)


func TestCreateShortURLWithAlias(t *testing.T) {

    models.InitDB()
    gin.SetMode(gin.TestMode) 

    r := routes.InitializeRoutes(gin.Default())


    // Assurez-vous que l'utilisateur est connecté ou simulez une authentification si nécessaire
    var jsonStr = []byte(`{"long_Url": "https://www.blogdumoderateur.com/tools/tinyurl/", "Alias": "exmpl"}`)
    req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }
    // Ajoutez des vérifications supplémentaires si nécessaire
}
