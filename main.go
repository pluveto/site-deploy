package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
	"github.com/pluveto/site-deploy/pkg/logger"
	"github.com/pluveto/site-deploy/pkg/setting"
	"github.com/pluveto/site-deploy/router"
)

func init() {
	logger.Setup()
	firstUse()
	setting.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := router.InitRouter()
	fmt.Printf("%+v\n", setting.ServerSetting)
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}

func firstUse() {
	_ = os.Mkdir("./tmp", os.ModeDir)
	_ = os.Mkdir("./log", os.ModeDir)
	_ = os.Mkdir("./conf", os.ModeDir)
	if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
		glg.Error("please configurate `conf/app.ini`")
		os.Exit(1)
	}
}
