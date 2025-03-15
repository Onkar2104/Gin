package main

import (
	"net/http"

	"github.com/Onkar2104/go/controllers"
	"github.com/Onkar2104/go/initializers"
	"github.com/Onkar2104/go/middleware"
	"github.com/Onkar2104/go/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {

	initializers.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	r.Static("/css", "./templates/css/")

	r.Static("/fonts", "./templates/fonts")

	r.Static("/img", "./templates/img/")

	r.Static("/js", "./templates/js")

	r.LoadHTMLGlob("templates/*.html")

	apiRoutes := r.Group("/post")
	{
		apiRoutes.POST("/register", controllers.Signup)
		apiRoutes.POST("/login/", controllers.Login)
		apiRoutes.POST("/upload", controllers.UploadFile)
		apiRoutes.GET("/validate", middleware.RequireAuth, controllers.Validate)

		apiRoutes.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	viewRoutes := r.Group("/view")
	{
		viewRoutes.GET("/register", func(c *gin.Context) {
			c.HTML(http.StatusOK, "signup.html", nil)
		})

		viewRoutes.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", nil)
		})

		viewRoutes.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})

		viewRoutes.GET("/upload", func(c *gin.Context) {
			c.HTML(http.StatusOK, "upload.html", nil)
		})

		// viewRoutes.GET("/files", func(c *gin.Context) {
		// 	c.HTML(http.StatusOK, "upload.html", nil)
		// })
	}

	r.Run()
}
