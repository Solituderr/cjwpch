package model

// TODO: add crud interface here

// 新增短链接
func addCreateUrl(createUrl CreateURL) (uint, error) {
	err := DB.Create(&createUrl).Error
	if err != nil {
		return 0, err
	}
	return createUrl.Id, err
}

// 新增用户信息
func addUser(register Register) error {
	err := DB.Create(Login{Email: register.Email, Pwd: register.Pwd}).Error
	if err != nil {
		return err
	}
	return err
}

// 验证登录信息是否有效
func getLogin(login Login) (string, error) {
	err := DB.Where("email = ? and pwd = ?", login.Email, login.Pwd).Error
	if err != nil {
		return "fail", err
	}
	return "successful", nil
}

// 查询短链接的详细信息
func inquireURL(ids uint) (error, CreateURL) {
	err := DB.Where("id = ?", "ids").Error
	if err != nil {
		return err, CreateURL{}
	}
	var createUrl CreateURL
	err1 := DB.First(&createUrl, ids).Error
	if err1 != nil {
		return err1, CreateURL{}
	}
	return nil, createUrl
}

// 获取用户的所有的短链接
func getUserAllURL(emails string) (error, []CreateURL) {
	err := DB.Where("email = ?", "emails").Error
	if err != nil {
		return err, []CreateURL{}
	}
	var createUrl []CreateURL
	err1 := DB.Find(&createUrl, emails).Error
	if err1 != nil {
		return err1, []CreateURL{}
	}
	return nil, createUrl
}

// UpdateShortURL 更新短链接
func updateShortURL(ids uint, update UpdateURL) error {
	err := DB.Where("id = ?", "ids").Error
	if err != nil {
		return err
	}
	var newURLInfo CreateURL
	err1 := DB.First(&newURLInfo, ids).Error
	if err1 != nil {
		return err1
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
	err := DB.Where("id = ?", "ids").Error
	if err != nil {
		return err
	}
	err1 := DB.Delete(&CreateURL{}, ids).Error
	if err1 != nil {
		return err1
	}
	return nil
}

/*
// 添加注册用户信息
func addRegister(register Register) error {
	err := DB.Create(&register).Error
	if err != nil {
		return err
	}
	return err
}

// 添加用户信息
func addLogin(login Login) error {
	err := DB.Create(&login).Error
	if err != nil {
		return err
	}
	return err
}
*/

// 获取用户信息
/*
func getUserInfo(register Register) (error, ReturnInfo) {
	err := DB.Model(&register).Where("email = ? and pwd = ?", register.Email, register.Pwd).Error
	if err != nil {
		return err, ReturnInfo{}
	}
	var returnInfo ReturnInfo
	returnInfo.Name = register.Name
	returnInfo.Email = register.Email
	return nil, returnInfo
}
*/
