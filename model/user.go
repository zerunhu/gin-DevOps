package model

type User struct {
	GdoModel
	Username  string   `gorm:"size:32;not null;unique comment:'用户名'"`
	Password  string   `gorm:"comment:'用户密码'"`
	Phone     string   `gorm:"comment:'电话'"`
	Email     string   `gorm:"comment:'邮箱'"`
	Group     []Group  `gorm:"many2many:app_user_group;"`
}
func (User) TableName() string {
	return "app_user"
}

type Group struct {
	GdoModel
	Name string `gorm:"size:255;not null;unique;comment:'组名'"`
	Desc string `gorm:"size:255;comment:'组名'"`
}
func (Group) TableName() string {
	return "app_group"
}

