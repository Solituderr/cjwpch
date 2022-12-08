package UserControl

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-svc-tpl/app/Response"
	"go-svc-tpl/model"
	"net/http"
)

type H map[string]interface{}

// Register 用户注册
func Register(c echo.Context) error {
	var register model.Register
	if err := c.Bind(&register); err != nil {
		logrus.Error("read register_information error!")
		return c.String(http.StatusBadRequest, "error")
	}
	//写进数据库接口

	//读取数据库目前id
	//id := LinkControl
	return Response.SenRes(c, http.StatusOK, "success", "")
}

// Login 用户登录
func Login(c echo.Context) error {
	var login model.Login
	if err := c.Bind(&login); err != nil {
		logrus.Error("login error!")
		return c.String(http.StatusBadRequest, "error")
	}
	//传给数据库数据
	isSuccess, err := "success", nil
	if isSuccess == "success" {
		return Response.SenRes(c,http.StatusOK,"success","")
	}
	else {
		return Response.SenRes(c,http.StatusBadRequest,"password or email is not right","")
	}
}

// LogOut 用户注销
func LogOut(c echo.Context) error {
	return nil
}

// Info 用户信息获取 感觉要id啊，不然怎么查
func Info(c echo.Context) error {
	//数据库读取
	name ,email ,err :=
	info := struct {
		name string
		email string
	}{name,email}
	if err!=nil{
		msg := "get info error"
		return Response.SenRes(c,http.StatusBadRequest,msg,"")
	} else{
		msg := "success"
		return Response.SenRes(c,http.StatusOK,msg,info)
	}
}


func GetAllUrl(c echo.Context) error {

}
