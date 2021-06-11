package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID        uint
	Username  string `gorm:"size:32;not null;unique;comment:'用户名'"`
	Groups     []Group `gorm:"many2many:user_group;"`
}
func (User) TableName() string {
	return "test_user"
}

type Group struct {
	ID        uint
	Name string `gorm:"size:32 comment:'组名'"`
}
func (Group) TableName() string {
	return "test_group"
}

//type UserGroup struct {
//	ID      uint
//	UserId  uint  `gorm:"ForeignKey:UserID"`
//	GroupId uint  `gorm:"ForeignKey:GroupID"`
//}
//func (UserGroup) TableName() string {
//	return "test_user_group"
//}


func main() {
	db, err := gorm.Open("mysql", "root:root123@tcp(172.20.1.167:3306)/test1?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	db.DropTable(&User{},&Group{})
	err = db.AutoMigrate(&User{},&Group{}).Error
	if err != nil{
		fmt.Println(err)
	}

	//a := Group{
	//	Name:  "admin",
	//}
	//db.Save(&a)
	//b := User{
	//	Username:  "ly",
	//}
	//db.Save(&b)

//err = db.Model(&user).Association("groups").Append(&group).Error
//err = db.Model(&user).Association("groups").Append(&groups).Error

//err = db.Model(&user).Association("groups").Delete(&groups).Error
//db.Model(&user).Association("groups").Delete(&group)

//err = db.Model(&user).Association("groups").Find(&groups).Error

//err = db.Model(&user).Association("groups").Replace(&group).Error 改是把原来的换成只有里面的
//err = db.Model(&user).Association("groups").Replace(&groups).Error
    //var user User
	//var group1 Group
	//db.First(&user)
	//db.First(&group1)
	//err = db.Model(&user).Association("groups").Replace(&group1).Error
	//if err != nil{
	//	fmt.Println(err)
	//}

	defer db.Close()
}
