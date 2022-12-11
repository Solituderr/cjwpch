package model

// TODO: add crud interface here

// 新增短链接（未登录）
func addLink(createUrl CreateURL) (uint, error) {
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

// 新增短链接（已登录）
func addLinkLogin(createUrl CreateURL, ids uint) (uint, error) {
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
	err1 := DB.Where("id = ?", "ids").First(&user).Error
	if err1 != nil {
		return 0, err
	}
	user.UrlId = append(user.UrlId, link.Id)
	return link.Id, nil
}

// 新增用户信息
func addUser(user User) error {
	err := DB.Create(&user).Error
	if err != nil {
		return err
	}
	return err
}

// 验证登录信息是否有效
func getLogin(login Login) (string, error) {
	var register User
	err := DB.Where("email = ? and pwd = ?", login.Email, login.Pwd).First(&register).Error
	if err != nil {
		return "fail", err
	}
	return "success", nil
}

// 查询短链接的详细信息
func inquireURL(ids uint) (error, Link) {
	var link Link
	err := DB.Where("id = ?", "ids").First(&link).Error
	if err != nil {
		return err, Link{}
	}
	return nil, link
}

// 获取用户的所有的短链接
func getUserAllURL(emails string) (error, []Link) {
	var link []Link
	err := DB.Where("email = ?", "emails").Find(&link).Error
	if err != nil {
		return err, []Link{}
	}
	return nil, link
}

// UpdateShortURL 更新短链接
func updateShortURL(ids uint, update UpdateURL) error {
	var newURLInfo Link
	err := DB.Where("id = ?", "ids").First(&newURLInfo).Error
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

// 删除短链接
func deleteShortUrl(ids uint) error {
	var link Link
	err := DB.Where("id = ?", "ids").Delete(&link).Error
	if err != nil {
		return err
	}
	return nil
}

// 暂停短链接
func pauseUrl(ids uint) error {
	var link Link
	err := DB.Where("id = ?", "ids").First(&link).Error
	if err != nil {
		return err
	}
	link.ExpireTime = link.StartTime
	DB.Save(&link)
	return nil
}
