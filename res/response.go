package res

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao9878/gin-common"
	"net/http"
)

type ResponseData struct {
	Code gin_common.ResCode `json:"code"`
	Msg  interface{}        `json:"msg"`
	Data interface{}        `json:"data"`
}

func Error(c *gin.Context, code gin_common.ResCode) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ErrorWithMsg(c *gin.Context, code gin_common.ResCode, msg interface{}) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func Success(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: gin_common.CodeSuccess,
		Msg:  gin_common.CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}
