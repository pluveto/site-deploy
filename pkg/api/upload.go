package api

import (
	"net/http"
	"path"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/pluveto/site-deploy/pkg/app"
	"github.com/pluveto/site-deploy/pkg/service"
	"github.com/pluveto/site-deploy/pkg/setting"
)

func Upload(c *gin.Context) {
	wrapper := app.Gin{C: c}
	file, err := c.FormFile("file")
	if err != nil {
		wrapper.Response(http.StatusBadRequest, 40000, "No file is received", nil)
		return
	}
	newFileName, err := randFileName(".zip")
	if err != nil {
		wrapper.Response(http.StatusInternalServerError, 50000, "What the f**k", nil)
		return
	}
	savePath := path.Join(setting.AppSetting.TempPath, newFileName)
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		wrapper.Response(http.StatusInternalServerError, 50000, "Unable to save the file", nil)
		return
	}

	err = service.Deploy(savePath, setting.AppSetting.SitePath)
	if err != nil {
		wrapper.Response(http.StatusInternalServerError, 50000, err.Error(), nil)
		return
	}
	wrapper.Response(http.StatusOK, 20000, "", nil)
}

func NoMethod(c *gin.Context) {
	wrapper := app.Gin{C: c}

	wrapper.Response(http.StatusBadRequest, 40000, "Method not allowed", nil)

}

func NotFound(c *gin.Context) {

	wrapper := app.Gin{C: c}

	wrapper.Response(http.StatusNotFound, 40000, "Not found", nil)
}

func randFileName(ext string) (string, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return "", err
	}
	return node.Generate().String() + ext, nil
}
