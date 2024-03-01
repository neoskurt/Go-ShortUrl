package routes

import (
	"go-url-shortener/controllers"

	"github.com/gin-gonic/gin"

    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
)

func InitializeRoutes(router *gin.Engine)  *gin.Engine{


        router = gin.Default()

        store := cookie.NewStore([]byte("secret"))
        router.Use(sessions.Sessions("mysession", store))

        router.Static("/static", "./static")
        router.LoadHTMLGlob("templates/*")

      // Route to serve the HTML file
      router.GET("/register", func(c *gin.Context) {
        c.HTML(200, "register.html", nil)
    })

      router.GET("/login", func(c *gin.Context) {
        c.HTML(200, "login.html", nil) 
    })

      router.GET("/", func(c *gin.Context) {
            session := sessions.Default(c)
            username := session.Get("username")
        if username == nil {
            c.HTML(200, "index.html", nil)
        } else {
            c.HTML(200, "index.html", gin.H{"Username": username})
        }
    })

    router.POST("/shorten", controllers.CreateShortURL)
    router.GET("/:short_url", controllers.RedirectShortURL)
    router.GET("/api/:short_url/stats", controllers.GetURLStatistics)
    router.POST("/register", controllers.RegisterUser)
    router.POST("/login", controllers.LoginUser)

    router.Use(func(c *gin.Context) {
        c.JSON(404, gin.H{"error": "Not Found"})
    })

    router.NoRoute(func(c *gin.Context) {
        c.JSON(404, gin.H{"error": "Not Found"})
    })

    return router
}
