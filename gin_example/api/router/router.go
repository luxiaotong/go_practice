package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//InitRouter is a function to Init Router
//参数:
//	无
//返回值:
//	*gin.Engine
func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"welcome": "Welcome! Thanks for using gintool",
		})
	})

	router.Static("/assets", "./assets")

	return router
}
