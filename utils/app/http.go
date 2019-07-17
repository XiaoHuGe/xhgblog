package app

import (
	"github.com/gin-gonic/gin"
)

// Response 团队基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   interface{} `json:"msg"`
	Error string      `json:"error"`
}

type Gin struct {
	C *gin.Context
}

func (this *Gin) Response(httpCode int, data *Response) {
	this.C.JSON(httpCode, data)
}
