package service

import (
	"errors"
	"fmt"
	"gin-DevOps/config"
	"gin-DevOps/model"
	"gin-DevOps/model/operate"
	"gin-DevOps/model/request"
	"gin-DevOps/model/response"
	"gin-DevOps/utils"
)

var (
	err error
)

func Login(user request.Login) error{
	sqlUser, err := operate.GetUserObjByField("username", user.Username)
	fmt.Println(sqlUser, err)
	if err != nil{
		return err
	}
	password := utils.MD5V([]byte(user.Password))
	if password != sqlUser.Password {
		return errors.New("password error")
	}
	return nil
}

func GetUserInfo(name string) ([]string, error){
	user, err := operate.GetUserObjByField("username", name)
	if err != nil{
		return nil, err
	}
	permissions, err := operate.GetPermissionsNameByUser(&user)
	return permissions, err
}

func ListUser() ([]response.UserListResponse, error){
	var users []model.User
	err = config.GdoDb.Find(&users).Error
	if err != nil{
		return nil, err
	}
	responseUser := make([]response.UserListResponse, 0, len(users))
	for _, user := range users{
		groups, err := operate.GetGroupsNameByUser(&user)
		if err != nil{
			return nil, err
		}
		u := response.UserListResponse{
			Id:    user.ID,
			Name:  user.Username,
			Email: user.Email,
			Phone: user.Phone,
			Group: groups,
		}
		responseUser = append(responseUser, u)
	}
	return responseUser, nil
}

func CreateUser(u request.User) error{
	_, err := operate.GetUserObjByField("username", u.Username)
	if err == nil{
		return errors.New("user name has been registered")
	}
	user := &model.User{
		Username: u.Username,
		Phone:    u.Phone,
		Email:    u.Email,
	}
	user.Password = utils.MD5V([]byte(u.Password))
	err = config.GdoDb.Create(&user).Error
	return err
}

func UpdateUser(userId string, u request.User) error {
	user := model.User{
		Username: u.Username,
		Phone:    u.Phone,
		Email:    u.Email,
	}
	user.Password = utils.MD5V([]byte(u.Password))
	err = config.GdoDb.Model(&user).Where("id = ?", userId).Updates(user).Error
	return err
}

func DeleteUser(userId string) error {
	var user model.User
	err = config.GdoDb.Where("id = ?", userId).Delete(&user).Error
	return err
}


//Group
func CreateGroup(g request.Group) error{
	_, err := operate.GetGroupObjByField("name", g.Name)
	if err == nil{
		return errors.New("group already exists")
	}
	group := &model.Group{
		Name:     g.Name,
		Desc:     g.Desc,
	}
	err = config.GdoDb.Create(&group).Error
	return err
}

func ListGroup() ([]model.Group, error) {
	var groups []model.Group
	err = config.GdoDb.Find(&groups).Error
	return groups, err
}

func UpdateGroup(groupId string, group model.Group) error {
	err = config.GdoDb.Model(&group).Where("id = ?", groupId).Updates(group).Error
	return err
}

func DeleteGroup(groupId string) error {
	var group model.Group
	err = config.GdoDb.Where("id = ?", groupId).Delete(&group).Error
	return err
}

//GroupUser
func ListGroupUser(groupId string) ([]response.GroupUserListResponse, error){
	users, err := operate.GetUserObjByGroup(groupId)
	if err != nil{
		return nil, err
	}
	responseUsers := make([]response.GroupUserListResponse, 0, len(users))
	for _, user := range users{
		responseUsers = append(responseUsers, response.GroupUserListResponse{
			Id:    user.ID,
			Name:  user.Username,
			Email: user.Email,
			Phone: user.Phone,
		})
	}
	return responseUsers, nil
}

func CreateGroupUser(userId string, groupId string) error{
	var user model.User
	var group model.Group
	user, err = operate.GetUserObjByField("id", userId)
	if err != nil{
		return err
	}
	group, err = operate.GetGroupObjByField("id", groupId)
	if err != nil{
		return err
	}
	if err = config.GdoDb.Model(&user).Association("Group").Append(&group); err != nil{
		return fmt.Errorf("group add user err: %w", err)
	}
	return nil
}

func DeleteGroupUser(userId []string, groupId string) error{
	var users []model.User
	var group model.Group
	err = config.GdoDb.Where("id IN ?", userId).Find(&users).Error
	if err != nil{
		return err
	}
	group, err = operate.GetGroupObjByField("id", groupId)
	if err != nil{
		return err
	}
	if err = config.GdoDb.Model(&group).Association("User").Delete(&users); err != nil{
		return fmt.Errorf("group delete user err: %w", err)
	}
	return nil
}

func ListPermission() ([]model.Permission, error) {
	var permissions []model.Permission
	err = config.GdoDb.Find(&permissions).Error
	return permissions, err
}
func CreatePermission(p request.Permission) error{
	_, err := operate.GetPermissionObjByField("name", p.Name)
	if err == nil{
		return errors.New("permission already exists")
	}
	permission := &model.Permission{
		Name:     p.Name,
		Desc:     p.Desc,
		Api:      p.Api,
		Action:   p.Action,
		Type:     p.Type,
	}
	err = config.GdoDb.Create(&permission).Error
	return err
}

func ListGroupPermission(groupId string) ([]model.Permission, error){
	group, err := operate.GetGroupObjByField("id", groupId)
	if err != nil{
		return nil, err
	}
	var permissions []model.Permission
	err = config.GdoDb.Model(&group).Association("Permission").Find(&permissions)
	if err != nil{
		return permissions, nil
	}
	return permissions, nil
}

func CreateGroupPermission(permissionId string, groupId string) error{
	permission, err := operate.GetPermissionObjByField("id", permissionId)
	if err != nil{
		return err
	}
	group, err := operate.GetGroupObjByField("id", groupId)
	if err != nil{
		return err
	}
	err = config.GdoDb.Model(&group).Association("Permission").Append(&permission)
	if err != nil{
		return fmt.Errorf("group add user err: %w", err)
	}
	return nil
}

func DeleteGroupPermission(permissionId []string, groupId string) error{
	var permissions []model.Permission
	var group model.Group
	err = config.GdoDb.Where("id IN ?", permissionId).Find(&permissions).Error
	if err != nil{
		return err
	}
	group, err = operate.GetGroupObjByField("id", groupId)
	if err != nil{
		return err
	}
	if err = config.GdoDb.Model(&group).Association("Permission").Delete(&permissions); err != nil{
		return fmt.Errorf("group delete user err: %w", err)
	}
	return nil
}