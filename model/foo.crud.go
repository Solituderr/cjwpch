package model

// TODO: add crud interface here

func addCreateUrl(createUrl CreateURL) (uint, error) {
	err := DB.Create(&createUrl).Error
	if err != nil {
		return 0, err
	}
	return createUrl.Id, err
}

/*
func addRegister(register Register) error {
	err := DB.Create(&register).Error
	if err != nil {
		return err
	}
	return err
}
*/

func addLogin(login Login) error {
	err := DB.Create(&login).Error
	if err != nil {
		return err
	}
	return err
}
