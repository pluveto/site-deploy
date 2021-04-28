package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pluveto/site-deploy/pkg/api"
	"github.com/pluveto/site-deploy/pkg/middleware"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.ValidateKey())
	r.POST("/upload", api.Upload)
	r.NoMethod(api.NoMethod)
	r.NoRoute(api.NotFound)
	return r
}
