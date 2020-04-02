package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luxiaotong/go_practice/gin_example/api/apis"
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

	router.GET("/getDatabaseList", apis.GetDatabaseList)
	router.GET("/getTableList", apis.GetTableList)
	router.GET("/getColumnList", apis.GetColumnList)

	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"welcome": "Welcome! Thanks for using gintool",
		})
	})

	router.Static("/assets", "./assets")

	return router
}
