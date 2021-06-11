package model

type User struct {
	GdoModel
	Username  string `gorm:"size:255 comment:'用户名'"`
	Password  string `gorm:"comment:'用户密码'"`
	Phone     string `gorm:"comment:'电话'"`
	Email     string `gorm:"comment:'邮箱'"`
}
func (User) TableName() string {
	return "app_user"
}
