package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	log.Println("start web server")
	return gin.Default()
}
