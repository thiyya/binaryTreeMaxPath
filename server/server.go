package server

import (
	"binaryTreeMaxPath/handlers"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	router := gin.Default()
	router.POST("/max-path", func(ctx *gin.Context) {
		handlers.HandleBinaryTreeMaxPath(ctx)
	})
	return router
}
