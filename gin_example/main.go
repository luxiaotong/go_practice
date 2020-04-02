package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itkinside/itkconfig"
	"net/http"
	"time"
	"fmt"
)

var AppConfig *AppConfigEntity

func main() {
	AppConfig = initConf()
	fmt.Printf("config : %s\n", AppConfig.HTTP_PORT)

	router := gin.Default()

	router.Static("/assets", "./assets")

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

	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", AppConfig.HTTP_HOST, AppConfig.HTTP_PORT),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func initConf() *AppConfigEntity {
	AppConfig := &AppConfigEntity{

	}
	itkconfig.LoadConfig("app.conf", AppConfig)
	if AppConfig.DEBUG == "false" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	fmt.Printf("config is : %s\n", AppConfig.DEBUG)
	return AppConfig
}

type AppConfigEntity struct {
	HTTP_HOST string
	HTTP_PORT string
	DEBUG     string
	APP_NAME  string
}
