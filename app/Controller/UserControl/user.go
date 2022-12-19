package UserControl

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-svc-tpl/app/Response"
	"go-svc-tpl/model"
	"go-svc-tpl/service"
	"net/http"
	"time"
)

var crud service.CRUD = service.Deal{}

type H map[string]interface{}

// Register 用户注册
func Register(c echo.Context) error {
	var register model.Register
	if err := c.Bind(&register); err != nil {
		logrus.Error("read register_information error!")
		return Response.SenRes(c, 400, "fail")
	}
	//写进数据库接口
	if err := crud.AddUser(register); err != nil {
		logrus.Error("add user error")
		return Response.SenRes(c, 400, "fail")
	} else {
		//读取数据库目前id
		//id := LinkControl
		return Response.SenRes(c, 200, "success")
	}
}

// Login 用户登录
func Login(c echo.Context) error {
	var login model.Login
	if err := c.Bind(&login); err != nil {
		logrus.Error("login error!")
		return Response.SenRes(c, 400, "fail")
	}
	//传给数据库数据
	isSuccess, id, err := crud.CheckUser(login)
	if err != nil {
		logrus.Error("check user error!")
		return Response.SenRes(c, 400, err.Error())
	}
	if isSuccess == "fail" {
		return Response.SenRes(c, 400, "password or email is not right")
	} else {
		claims := &model.JwtCustomClaims{
			Id:   id,
			Name: login.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return Response.SenRes(c, 200, "success", echo.Map{
			"token": t,
		})
	}
}

// LogOut 用户注销
func LogOut(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	claims.ExpiresAt = time.Now().Unix()
	return Response.SenRes(c, http.StatusOK, "log out successful")
}

// Info 用户信息获取
func Info(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	s := claims.Name
	//数据库读取
	name, pwd, err := crud.GetInfoUser(s)
	info := struct {
		name string
		pwd  string
	}{name, pwd}
	if err != nil {
		msg := "get info error"
		return Response.SenRes(c, http.StatusBadRequest, msg, "")
	} else {
		msg := "success"
		return Response.SenRes(c, http.StatusOK, msg, info)
	}
}

func GetAllUrl(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	s := claims.Name //EMAIL
	if urls, err := crud.GetUrlsUser(s); err != nil {
		logrus.Error("get urls error!")
		return Response.SenRes(c, 400, "fail")
	} else {
		return Response.SenRes(c, 200, "success", urls)
	}
}
