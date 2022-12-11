package model

// TODO: add new model

import (
	. "time"
)

// User 用户注册信息
type User struct {
	Id    uint   `json:"id" gorm:"primaryKey" form:"id"`
	Name  string `json:"name"  form:"name"`
	Email string `json:"email" form:"email"`
	Pwd   string `json:"pwd"   form:"pwd"`
	UrlId []uint `json:"urlId" form:"urlId"`
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
	Id         uint   `json:"id" gorm:"primaryKey" form:"id"`
	Origin     string `json:"origin"               form:"origin"`
	Short      string `json:"short"                form:"short"`
	Comment    string `json:"comment"              form:"comment"`
	StartTime  Time   `json:"startTime"            form:"startTime"`
	ExpireTime Time   `json:"expireTime"           form:"expireTime"`
}

type CreateURL struct {
	Origin     string `json:"origin"               form:"origin"`
	Short      string `json:"short"                form:"short"`
	Comment    string `json:"comment"              form:"comment"`
	StartTime  Time   `json:"startTime"            form:"startTime"`
	ExpireTime Time   `json:"expireTime"           form:"expireTime"`
}

// UpdateURL 更新短链接
type UpdateURL struct {
	Id         uint   `json:"id" gorm:"primaryKey" form:"id"`
	Origin     string `json:"origin"               form:"origin"`
	Comment    string `json:"comment"              form:"comment"`
	StartTime  Time   `json:"startTime"            form:"startTime"`
	ExpireTime Time   `json:"expireTime"           form:"expireTime"`
}

// ReturnInfo 返回用户信息
type ReturnInfo struct {
	Name  string `json:"name"  form:"name"`
	Email string `json:"email" form:"email"`
}
