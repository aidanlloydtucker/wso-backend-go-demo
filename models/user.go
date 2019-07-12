package models

//UserModel ...
type UserModel struct {
	BaseModel
}

func (m *UserModel) GetAllUsers(u *[]User) (err error) {
	if err = m.DB.Set("gorm:auto_preload", true).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func (m *UserModel) GetUserByID(id uint, u *User) (err error) {
	if err = m.DB.Set("gorm:auto_preload", true).Where(NewUserWithID(id)).First(u).Error; err != nil {
		return err
	}
	return nil
}
