package Response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Date interface{} `json:"date"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = 200
	ERROR   = 500
)

func Result(code int, date interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Date: date,
		Msg:  msg,
	})
}
func ResponseSuccess(date interface{}, mes string, c *gin.Context) {
	Result(SUCCESS, date, mes, c)
}
func ResponseRrror(date interface{}, mes string, c *gin.Context) {
	Result(ERROR, date, mes, c)
}
