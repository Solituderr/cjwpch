package response

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

// Response 统一的response
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// SendResponse 发送回应
func SendResponse(c echo.Context, code int, msg string, data interface{}) error {

	err3 := c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	if err3 != nil {
		log.Fatal(err3)
	}
	return nil
}
