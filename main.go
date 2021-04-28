package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pluveto/site-deploy/pkg/logger"
	"github.com/pluveto/site-deploy/pkg/setting"
	"github.com/pluveto/site-deploy/router"
)

// 单文件上传
func fileUpload(context *gin.Context) {
	file, err := context.FormFile("file")
	if err != nil {
		log.Println("ERROR: upload file failed. ", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Sprintf("ERROR: upload file failed. %s", err),
		})
	}
	dst := "./tmp/" + file.Filename

	err = context.SaveUploadedFile(file, dst)
	if err != nil {
		log.Println("ERROR: save file failed. ", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Sprintf("ERROR: save file failed. %s", err),
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"date":    nil,
		"message": nil,
	})
}

func init() {
	setting.Setup()
	logger.Setup()
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
