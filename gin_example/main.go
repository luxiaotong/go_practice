package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itkinside/itkconfig"
	"github.com/luxiaotong/go_practice/gin_example/api/database"
	"github.com/luxiaotong/go_practice/gin_example/api/router"
)

//AppConfig contains configs of application
var AppConfig *AppConfigEntity

func main() {
	AppConfig = initConf()
	fmt.Printf("config : %s\n", AppConfig.PORT)
	router := router.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", AppConfig.HOST, AppConfig.PORT),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		http.ListenAndServe("127.0.0.1:6060", nil)
	}()

	s.ListenAndServe()
	database.Eloquent.Close()
}

func initConf() *AppConfigEntity {
	AppConfig := &AppConfigEntity{}
	itkconfig.LoadConfig("app.conf", AppConfig)
	if AppConfig.DEBUG == "false" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	fmt.Printf("config is : %s\n", AppConfig.DEBUG)
	return AppConfig
}

// AppConfigEntity contains config
type AppConfigEntity struct {
	HOST    string
	PORT    string
	DEBUG   string
	APPNAME string
}
