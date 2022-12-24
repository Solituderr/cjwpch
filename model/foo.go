package model

// TODO: add new model

import (
	"time"
)

type Time time.Time

// User 用户注册信息
type User struct {
	Id      uint   `json:"id" gorm:"primaryKey;auto_increment" form:"id"`
	Name    string `json:"name"  form:"name"`
	Email   string `json:"email" form:"email"`
	Pwd     string `json:"pwd"   form:"pwd"`
	UrlInfo []Link `json:"urlInfo" gorm:"foreignKey:UserId" form:"urlInfo"`
}

// Register 用户注册信息
type Register struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Pwd   string `json:"pwd" form:"pwd"`
}

// Login 用户登录信息
type Login struct {
	Email string `json:"email" form:"email"`
	Pwd   string `json:"pwd"   form:"pwd"`
}

// Link CreateURL 短连接信息
type Link struct {
	Id         uint      `json:"id" gorm:"primaryKey;auto_increment" form:"id"`
	UserId     string    `json:"userId"               form:"userId"`
	Origin     string    `json:"origin"               form:"origin"`
	Short      string    `json:"short"                form:"short"`
	Comment    string    `json:"comment"              form:"comment"`
	StartTime  time.Time `json:"startTime"            form:"startTime"`
	ExpireTime time.Time `json:"expireTime"           form:"expireTime"`
}

type CreateURL struct {
	Origin     string    `json:"origin"               form:"origin"`
	Short      string    `json:"short"                form:"short"`
	Comment    string    `json:"comment"              form:"comment"`
	StartTime  time.Time `json:"startTime"            form:"startTime"`
	ExpireTime time.Time `json:"expireTime"           form:"expireTime"`
}

// UpdateURL 更新短链接
type UpdateURL struct {
	Id         uint      `json:"id" gorm:"primaryKey" form:"id"`
	Origin     string    `json:"origin"               form:"origin"`
	Comment    string    `json:"comment"              form:"comment"`
	StartTime  time.Time `json:"startTime"            form:"startTime"`
	ExpireTime time.Time `json:"expireTime"           form:"expireTime"`
}

// ReturnInfo 返回用户信息
type ReturnInfo struct {
	Name  string `json:"name"  form:"name"`
	Email string `json:"email" form:"email"`
}

//func (t *Time) UnmarshalParam(src string) error {
//	ts, err := time.Parse(time.RFC3339, src)
//	*t = Time(ts)
//	return err
//}
