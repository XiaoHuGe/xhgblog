package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHtml(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index/index.html", nil)
}
