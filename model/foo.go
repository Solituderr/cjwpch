package model

// TODO: add new model

import (
	. "time"
)

type Register struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}

type Login struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}

type CreateURL struct {
	Id         uint   `json:"id" gorm:"primaryKey"`
	Origin     string `json:"origin"`
	Short      string `json:"short"`
	Comment    string `json:"comment"`
	StartTime  Time   `json:"startTime"`
	ExpireTime Time   `json:"expireTime"`
}

type UpdateURL struct {
	Origin     string `json:"origin"`
	Comment    string `json:"comment"`
	StartTime  Time   `json:"startTime"`
	ExpireTime Time   `json:"expireTime"`
}
