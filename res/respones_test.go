package res

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao9878/gin-common"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	resCode gin_common.ResCode = http.StatusOK
	r       *gin.Engine
	res     = &ResponseData{
		Code: resCode,
		Msg:  resCode.Msg(),
		Data: nil,
	}
)

func init() {
	gin.DisableConsoleColor()
	r = gin.Default()
	r.GET("/test", func(c *gin.Context) {
		Success(c, nil)
	})
}

func TestSuccess(t *testing.T) {
	url := "/test"
	resp := string(Get(url, r))
	if resp != `{"code":1000,"msg":"success","data":null}` {
		t.Errorf("响应数据不符,res:%v", resp)
	}
}

// Get 根据特定请求uri，发起get请求返回响应
func Get(uri string, router *gin.Engine) []byte {
	// 构造get请求
	req := httptest.NewRequest("GET", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}
