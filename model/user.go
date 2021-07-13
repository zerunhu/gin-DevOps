package model

type User struct {
	GdoModel
	Username  string   `gorm:"size:32;not null;unique;comment:'用户名'"`
	Password  string   `gorm:"comment:'用户密码'"`
	Phone     string   `gorm:"comment:'电话'"`
	Email     string   `gorm:"comment:'邮箱'"`
	Group     []Group  `gorm:"many2many:auth_user_group;"`
}
func (User) TableName() string {
	return "auth_user"
}

type Group struct {
	GdoModel
	Name        string        `json:"name" gorm:"size:64;not null;unique;comment:'组名'"`
	Desc        string        `json:"desc" gorm:"size:64;comment:'组名'"`
	User        []User        `gorm:"many2many:auth_user_group;"`
	Permission  []Permission  `gorm:"many2many:auth_group_permission;"`
}
func (Group) TableName() string {
	return "auth_group"
}

type Permission struct {
	GdoModel
	Name    string  `json:"name" gorm:"size:64;not null;unique;comment:'组名'"`
	Desc    string  `json:"desc" gorm:"size:64;comment:'组名'"`
	Api     string  `json:"api" gorm:"size:128;comment:'接口'"`
	Action  string  `json:"action" gorm:"size:64;comment:'动作'"`
	Type    string  `json:"type" gorm:"size:64;comment:'类型'"`
	Group   []Group `gorm:"many2many:auth_group_permission;"`
}
func (Permission) TableName() string {
	return "auth_permission"
}

