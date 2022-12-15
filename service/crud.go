package service

import (
	"go-svc-tpl/model"
)

type LinkDeal interface {
	CreateUrl(model.CreateURL) (uint, error)
	CreateUrlLogin(model.CreateURL, uint) (uint, error) //id
	InquireUrl(uint) (model.CreateURL, error)
	UpdateUrl(model.UpdateURL) error
	DeleteUrl(uint) error
	PauseUrl(uint) error
	GetUrl(string) (string, error) //传短链接，返回重定向到的链接
}

type UserDeal interface {
	AddUser(model.Register) error          //创建用户，返回是否有错误
	CheckUser(model.Login) (string, error) //返回success或fail  判断登录是否成功
	LogOutUser() error
	GetInfoUser(string) (string, string, error)

	GetUrlsUser(string) ([]model.CreateURL, error)
}

type CRUD interface {
	LinkDeal
	UserDeal
}

type Deal struct{}

func (Deal) CreateUrl(url model.CreateURL) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (Deal) CreateUrlLogin(url model.CreateURL, id uint) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (Deal) InquireUrl(id uint) (model.CreateURL, error) {
	//TODO implement me
	panic("implement me")
}

func (Deal) UpdateUrl(link model.UpdateURL) error {
	//TODO implement me
	panic("implement me")
}

func (Deal) DeleteUrl(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (Deal) PauseUrl(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (Deal) GetUrl(s string) (string, error) {
	panic("123")
}

func (Deal) AddUser(register model.Register) error {
	//TODO implement me
	panic("implement me")
}

func (Deal) CheckUser(login model.Login) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (Deal) LogOutUser() error {
	//TODO implement me
	panic("implement me")
}

func (Deal) GetInfoUser(s string) (string, string, error) {
	//TODO implement me
	panic("implement me")
}

func (Deal) GetUrlsUser(s string) ([]model.CreateURL, error) {
	//TODO implement me
	panic("implement me")
}
