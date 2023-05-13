package server

import "github.com/gin-gonic/gin"

func NewHTTPServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.Default()
}
