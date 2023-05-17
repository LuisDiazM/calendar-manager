package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ApiKeyMiddleware(apiKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKeyValue := ctx.Request.Header.Get(API_KEY_NAME)
		if !strings.EqualFold(apiKeyValue, apiKey) {
			ctx.JSON(http.StatusUnauthorized, nil)
			ctx.Abort()
		}
		ctx.Next()
	}
}
