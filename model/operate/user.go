package operate

import (
	"errors"
	"gin-DevOps/config"
	"gin-DevOps/model"
)

var (
	err error
)


func GetUserObjByField(field string, name string) (model.User, error) {
	var user model.User
	err = config.GdoDb.Where(field + " = ?", name).First(&user).Error
	if err != nil{
		return user, errors.New("user does not exist")
	}
	return user, nil
}

func GetGroupObjByField(field string, name string) (model.Group, error) {
	var group model.Group
	err = config.GdoDb.Where(field + " = ?", name).First(&group).Error
	if err != nil{
		return group, errors.New("user does not exist")
	}
	return group, nil
}

func GetPermissionObjByField(field string, name string) (model.Permission, error) {
	var permission model.Permission
	err = config.GdoDb.Where(field + " = ?", name).First(&permission).Error
	if err != nil{
		return permission, errors.New("permission does not exist")
	}
	return permission, nil
}

func GetUserObjByGroup(groupId string) ([]model.User, error){
	group, err := GetGroupObjByField("id", groupId)
	if err != nil{
		return nil, err
	}
	var users []model.User
	err = config.GdoDb.Model(&group).Association("User").Find(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetGroupsObjByUser(user *model.User) ([]model.Group, error) {
	var groups []model.Group
	err = config.GdoDb.Model(user).Association("Group").Find(&groups)
	if err != nil{
		return nil, err
	}
	return groups, nil
}

func GetGroupsNameByUser(user *model.User) ([]string, error) {
	groups, err := GetGroupsObjByUser(user)
	if err != nil{
		return nil, err
	}
	groupsName := make([]string, 0, len(groups))
	for _, group := range groups{
		groupsName = append(groupsName, group.Name)
	}
	return groupsName, nil
}


func GetPermissionsObjByGroup(group *model.Group) ([]model.Permission, error){
	var permissions []model.Permission
	err = config.GdoDb.Model(group).Association("Permission").Find(&permissions)
	if err != nil{
		return nil, err
	}
	return permissions, nil
}

func GetPermissionsNameByUser(user *model.User) ([]string, error){
	groups, err := GetGroupsObjByUser(user)
	if err != nil {
		return nil, err
	}
	permissionMap := make(map[string]struct{})
	for _, group := range groups{
		pers, err := GetPermissionsObjByGroup(&group)
		if err != nil {
			return nil, err
		}
		for _, p := range pers{
			if _, ok := permissionMap[p.Name]; !ok{
				permissionMap[p.Name] = struct{}{}
			}
		}
	}
	permissions := make([]string, 0, len(permissionMap))
	for key, _ := range permissionMap{
		permissions = append(permissions, key)
	}
	return permissions, nil
}