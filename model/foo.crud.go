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

func addUser(register Register) error {
	err := DB.Create(Login{Email: register.Email, Pwd: register.Pwd}).Error
	if err != nil {
		return err
	}
	return err
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

// 验证登录信息是否有效
func getLogin(login Login) (string, error) {
	err := DB.Where("email = ? and pwd = ?", login.Email, login.Pwd).Error
	if err != nil {
		return "fail", err
	}
	return "successful", nil
}

// 获取用户信息
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
