package model

import "time"

// TODO: add crud interface here

// AddLink 新增短链接（未登录）
func AddLink(createUrl CreateURL) (uint, error) {
	var link Link
	link.Origin = createUrl.Origin
	link.Short = createUrl.Short
	link.Comment = createUrl.Comment
	link.StartTime = createUrl.StartTime
	link.ExpireTime = createUrl.ExpireTime
	err := DB.Create(&link).Error
	if err != nil {
		return 0, err
	}
	return link.Id, nil
}

// AddLinkLogin 新增短链接（已登录）
func AddLinkLogin(createUrl CreateURL, ids uint) (uint, error) {
	var link Link
	var user User
	link.Origin = createUrl.Origin
	link.Short = createUrl.Short
	link.Comment = createUrl.Comment
	link.StartTime = createUrl.StartTime
	link.ExpireTime = createUrl.ExpireTime
	err := DB.Create(&link).Error
	if err != nil {
		return 0, err
	}
	err1 := DB.Where("id = ?", ids).First(&user).Error
	if err1 != nil {
		return 0, err
	}
	user.UrlInfo = append(user.UrlInfo, link)
	return link.Id, nil
}

// InquireURL 查询短链接的详细信息
func InquireURL(ids uint) (error, Link) {
	var link Link
	err := DB.Where("id = ?", ids).First(&link).Error
	if err != nil {
		return err, Link{}
	}
	return nil, link
}

// UpdateShortURL 更新短链接
func UpdateShortURL(ids uint, update UpdateURL) error {
	var newURLInfo Link
	err := DB.Where("id = ?", ids).First(&newURLInfo).Error
	if err != nil {
		return err
	}
	newURLInfo.Origin = update.Origin
	newURLInfo.Comment = update.Comment
	newURLInfo.StartTime = update.StartTime
	newURLInfo.ExpireTime = update.ExpireTime
	DB.Save(&newURLInfo)
	return nil
}

// DeleteShortUrl 删除短链接
func DeleteShortUrl(ids uint) error {
	var link Link
	err := DB.Where("id = ?", ids).Delete(&link).Error
	if err != nil {
		return err
	}
	return nil
}

// PauseUrl 暂停短链接
func PauseUrl(ids uint) error {
	var link Link
	err := DB.Where("id = ?", ids).First(&link).Error
	if err != nil {
		return err
	}
	link.ExpireTime = link.StartTime
	DB.Save(&link)
	return nil
}

// GetUrl 重定向短链接
func GetUrl(shortUrl string) (string, error) {
	var link Link
	err := DB.Where("short = ?", shortUrl).First(&link).Error
	if err != nil {
		return "fail", err
	}
	now := time.Now()
	if now.Before(link.ExpireTime) {
		return link.Origin, nil
	} else {
		return "expired", nil
	}
}

//-----------------------------------------------------//

// AddUser 新增用户信息
func AddUser(userInfo Register) error {
	var user User
	user.Name = userInfo.Name
	user.Pwd = userInfo.Pwd
	user.Email = userInfo.Email
	user.UrlInfo = []Link{}
	err := DB.Create(&user).Error
	if err != nil {
		return err
	}
	return err
}

// GetLogin 验证登录信息是否有效
func GetLogin(login Login) (string, uint, error) {
	var register User
	err := DB.Where("email = ? and pwd = ?", login.Email, login.Pwd).First(&register).Error
	if err != nil {
		return "fail", 0, err
	}
	return "success", register.Id, nil
}

// GetInfoUser 获取用户信息
func GetInfoUser(emails string) (string, string, error) {
	var user User
	err := DB.Where("email = ?", emails).Find(&user).Error
	if err != nil {
		return "", "", err
	}
	return user.Name, user.Pwd, nil
}

// GetUserAllURL 获取用户的所有的短链接
func GetUserAllURL(emails string) (error, []Link) {
	var user User
	err := DB.Where("email = ?", emails).Find(&user).Error
	if err != nil {
		return err, []Link{}
	}
	var link Link
	var links []Link
	for _, info := range user.UrlInfo {
		err1 := DB.Where("Id = ?", info.Id).Find(&link).Error
		if err1 != nil {
			continue
		}
		links = append(links, link)
	}
	return nil, links
}
