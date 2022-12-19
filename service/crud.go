package service

import (
	"go-svc-tpl/model"
)

type LinkDeal interface {
	CreateUrl(model.CreateURL) (uint, error)
	CreateUrlLogin(model.CreateURL, uint) (uint, error) //id
	InquireUrl(uint) (model.Link, error)
	UpdateUrl(uint, model.UpdateURL) error
	DeleteUrl(uint) error
	PauseUrl(uint) error
	GetUrl(string) (string, error) //传短链接，返回重定向到的链接
}

type UserDeal interface {
	AddUser(model.Register) error                //创建用户，返回是否有错误
	CheckUser(model.Login) (string, uint, error) //返回success或fail  判断登录是否成功
	GetInfoUser(string) (string, string, error)  //传email 返回name和pwd
	GetUrlsUser(string) ([]model.Link, error)
}

type CRUD interface {
	LinkDeal
	UserDeal
}

type Deal struct{}

func (Deal) CreateUrl(url model.CreateURL) (uint, error) {
	id, err := model.AddLink(url)
	return id, err
}

func (Deal) CreateUrlLogin(url model.CreateURL, id uint) (uint, error) {
	loginId, err := model.AddLinkLogin(url, id)
	return loginId, err
}

func (Deal) InquireUrl(id uint) (model.Link, error) {
	err, link := model.InquireURL(id)
	return link, err
}

func (Deal) UpdateUrl(id uint, link model.UpdateURL) error {
	err := model.UpdateShortURL(id, link)
	return err
}

func (Deal) DeleteUrl(id uint) error {
	err := model.DeleteShortUrl(id)
	return err
}

func (Deal) PauseUrl(id uint) error {
	err := model.PauseUrl(id)
	return err
}

func (Deal) GetUrl(s string) (string, error) {
	url, err := model.GetUrl(s)
	return url, err
}

//------------------------------------------------------------//

func (Deal) AddUser(register model.Register) error {
	err := model.AddUser(register)
	return err
}

func (Deal) CheckUser(login model.Login) (string, uint, error) {
	isSuccess, id, err := model.GetLogin(login)
	return isSuccess, id, err
}

func (Deal) GetInfoUser(s string) (string, string, error) {
	name, pwd, err := model.GetInfoUser(s)
	return name, pwd, err
}

func (Deal) GetUrlsUser(s string) ([]model.Link, error) {
	err, links := model.GetUserAllURL(s)
	return links, err
}
