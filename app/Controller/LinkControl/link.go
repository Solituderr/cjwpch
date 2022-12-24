package LinkControl

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-svc-tpl/app/Response"
	"go-svc-tpl/model"
	"go-svc-tpl/service"
	"net/http"
	"strconv"
)

type H map[string]interface{}

var crud service.CRUD = service.Deal{}

func CreateLink(c echo.Context) error {
	var create model.CreateURL
	if err := c.Bind(&create); err != nil {
		logrus.Error("bind createUrl error!")
		return Response.SenRes(c, 400, "fail")
	}
	create.Short, _ = service.Transform(create.Origin)
	//调用数据库接口返回是否创建成功
	if id, err := crud.CreateUrl(create); err != nil {
		logrus.Error("create link error!")
		return Response.SenRes(c, 400, "fail")
	} else {
		return Response.SenRes(c, 200, "success", echo.Map{
			"id": id,
		})
	}
}

func CreateLinkLogin(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	userId := uint(claims.Id)
	var create model.CreateURL
	if err := c.Bind(&create); err != nil {
		logrus.Error(err)
		return Response.SenRes(c, 400, "fail")
	}
	create.Short, _ = service.Transform(create.Origin)
	if id, err := crud.CreateUrlLogin(create, userId); err != nil {
		logrus.Error("create link error")
		return Response.SenRes(c, 400, "fail")
	} else {
		return Response.SenRes(c, 200, "success", echo.Map{
			"id": id,
		})
	}
}

func QueryLink(c echo.Context) error {
	id := c.FormValue("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error("string -> int error!")
	}
	//调用数据库接口返回信息
	if linkInfo, err := crud.InquireUrl(uint(i)); err != nil {
		logrus.Error(err)
		return Response.SenRes(c, 400, "fail")
	} else {
		return Response.SenRes(c, 200, "success", linkInfo)
	}
}

func UpdateLink(c echo.Context) error {
	var update model.UpdateURL
	if err := c.Bind(&update); err != nil {
		logrus.Error(err)
		return Response.SenRes(c, 400, "fail")
	}
	id := update.Id
	//调用数据库接口返回是否创建成功
	if err := crud.UpdateUrl(id, update); err != nil {
		logrus.Error(err)
		return Response.SenRes(c, 400, "fail")
	} else {
		return Response.SenRes(c, 200, "success")
	}
}

func DeleteLink(c echo.Context) error {
	id := c.FormValue("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error("string -> int error!")
	}
	//调用数据库接口返回是否删除成功
	if err = crud.DeleteUrl(uint(i)); err != nil {
		logrus.Error("delete url error!")
		return Response.SenRes(c, 400, "fail")
	} else {
		return Response.SenRes(c, 200, "success")
	}
}

func PauseLink(c echo.Context) error {
	id := c.FormValue("id")
	//调用数据库接口返回是否暂停成功
	i, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error("string -> int error!")
	}
	//调用数据库接口返回是否删除成功
	if err = crud.PauseUrl(uint(i)); err != nil {
		logrus.Error("pause url error!")
		return Response.SenRes(c, 400, "fail")
	} else {
		return Response.SenRes(c, 200, "success")
	}
}

func Redirect(c echo.Context) error {
	sl := c.Param("shortLink")
	url, err := crud.GetUrl(sl)
	if url == "expired" {
		return Response.SenRes(c, 400, "shortlink is expired")
	}
	if err != nil {
		return Response.SenRes(c, 400, "fail")
	} else {
		return c.Redirect(http.StatusFound, url)
	}
}
