package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoRouterHtml(c *gin.Context) {
	HandleMessage(c, "找不到页面")
}

func HandleMessage(c *gin.Context, message string) {
	c.HTML(http.StatusNotFound, "errors/error.html", gin.H{
		"message": message,
	})
}
