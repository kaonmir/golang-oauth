package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaonmir/OAuth/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
		})
	})
	router.GET("/naver", func(c *gin.Context) {
		c.HTML(http.StatusOK, "naver.html", gin.H{})
	})
	router.GET("/callback", func(c *gin.Context) {
		c.HTML(http.StatusOK, "callback.html", gin.H{})
	})

	// router.Use(middlewares.AuthMiddleware())

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			userGroup := v1.Group("user")
			{
				userGroup.POST("/", controllers.CreateUser)
				userGroup.GET("/", controllers.GetUserByID)

				callback := userGroup.Group("callback")
				{
					callback.GET("/naver", controllers.NaverCallBackHandler)
				}
			}
		}
	}
	return router

}
