package middleware

import (
	"net/http"
	"strings"

	"github.com/pluveto/site-deploy/pkg/app"
	"github.com/pluveto/site-deploy/pkg/setting"
	"github.com/gin-gonic/gin"
)

func ValidateKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		wrapper := app.Gin{C: c}

		token := c.GetHeader("authorization") //从请求的header中获取toekn字符串
		validKey := strings.TrimSpace(setting.AppSetting.Key)
		if len(validKey) < 32 {
			wrapper.Response(http.StatusForbidden, 40000, "Weak or empty key", validKey)
			c.Abort()
			return
		}

		key := last(strings.Split(token, " "))
		if strings.TrimSpace(key) != validKey {
			wrapper.Response(http.StatusForbidden, 40000, "No permission", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

func last(s []string) string {
	if len(s) == 0 {
		return ""
	}
	return s[len(s)-1]
}
