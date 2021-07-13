package api

import (
	"gin-DevOps/config"
	"gin-DevOps/middleware"
	"gin-DevOps/model"
	"gin-DevOps/model/request"
	"gin-DevOps/model/response"
	"gin-DevOps/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

var (
	err error
)

// LoginHandler 登录接口
// @Summary 登录接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags login相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "JWT 用户令牌"
// @Param data body model.User true "用户名, 用户密码"
// @Security ApiKeyAuth
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/login [post]
func Login(c *gin.Context) {
	var user request.Login
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithErrMessage("Login failed", err, c)
		return
	}
	err = service.Login(user)
	if err != nil {
		config.GdoLog.Error("Login failed", zap.Any("err",err))
		response.FailWithErrMessage("Login failed", err, c)
		return
	}
	tokenString, _ := middleware.GenToken(user.Username)
	data := map[string]string{
		"token": tokenString,
	}
	response.OkWithDetailed(data, "Login success", c)
}

func UserInfo(c *gin.Context) {
	username, _ := c.Get("username")
	permissions, err := service.GetUserInfo(username.(string))
	if err != nil{
		config.GdoLog.Error("Get userinfo failed", zap.Any("err",err))
		response.FailWithErrMessage("Get userinfo failed", err, c)
		return
	}
	data := map[string]interface{}{
		"username": username,
		"permissions": permissions,
	}
	response.OkWithDetailed(data, "Get userinfo success", c)
}

func ListUser(c *gin.Context) {
	users, err := service.ListUser()
	if err != nil{
		config.GdoLog.Error("List user failed", zap.Any("err",err))
		response.FailWithErrMessage("List user failed", err, c)
		return
	}
	response.OkWithDetailed(users, "List user success", c)
}

func CreateUser(c *gin.Context) {
	var user request.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithErrMessage("Create user failed", err, c)
		return
	}
	err = service.CreateUser(user)
	if err != nil {
		config.GdoLog.Error("Create user failed", zap.Any("err", err))
		response.FailWithErrMessage("Create user failed", err, c)
		return
	}
	response.OkWithMessage("Create user success", c)
}

func UpdateUser(c *gin.Context){
	userId := c.Param("uid")
	var user request.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithErrMessage("Update user failed", err, c)
		return
	}
	err = service.UpdateUser(userId, user)
	if err != nil {
		config.GdoLog.Error("Update user failed", zap.Any("err", err))
		response.FailWithErrMessage("Update user failed", err, c)
		return
	}
	response.OkWithMessage("Update user success", c)
}

func DeleteUser(c *gin.Context){
	userId := c.Param("uid")
	err = service.DeleteUser(userId)
	if err != nil {
		response.FailWithErrMessage("Delete user failed", err, c)
		return
	}
	response.OkWithMessage("Delete user success", c)
}


func CreateGroup(c *gin.Context){
	var group request.Group
	err = c.ShouldBindJSON(&group)
	if err != nil {
		response.FailWithErrMessage("Create group failed", err, c)
		return
	}
	err = service.CreateGroup(group)
	if err != nil {
		config.GdoLog.Error("Create group failed", zap.Any("err", err))
		response.FailWithErrMessage("Create group failed", err, c)
		return
	}
	response.OkWithMessage("Create group success", c)
}

func ListGroup(c *gin.Context) {
	groups, err := service.ListGroup()
	if err != nil{
		config.GdoLog.Error("List group failed", zap.Any("err", err))
		response.FailWithErrMessage("List group failed", err, c)
		return
	}
	response.OkWithDetailed(groups, "List group success", c)
}

func UpdateGroup(c *gin.Context){
	groupId := c.Param("gid")
	var group model.Group
	err = c.ShouldBindJSON(&group)
	if err != nil {
		response.FailWithErrMessage("Update group failed", err, c)
		return
	}
	err = service.UpdateGroup(groupId, group)
	if err != nil {
		config.GdoLog.Error("Update group failed", zap.Any("err", err))
		response.FailWithErrMessage("Update group failed", err, c)
		return
	}
	response.OkWithMessage("Update group success", c)
}

func DeleteGroup(c *gin.Context){
	groupId := c.Param("gid")
	err = service.DeleteGroup(groupId)
	if err != nil{
		config.GdoLog.Error("Delete group failed", zap.Any("err", err))
		response.FailWithErrMessage("Delete group failed", err, c)
		return
	}
	response.OkWithMessage("Delete group success", c)
}



func ListGroupUser(c *gin.Context){
	groupId := c.Param("gid")
	users, err := service.ListGroupUser(groupId)
	if err != nil{
		config.GdoLog.Error("List users of the group failed", zap.Any("err", err))
		response.FailWithErrMessage("List users of the group failed", err, c)
		return
	}
	response.OkWithDetailed(users, "List users of the group success", c)
}

func CreateGroupUser(c *gin.Context){
	userId := c.Param("uid")
	groupId := c.Param("gid")
	err = service.CreateGroupUser(userId, groupId)
	if err != nil{
		config.GdoLog.Error("Create users of the group failed", zap.Any("err", err))
		response.FailWithErrMessage("Create users of the group failed", err, c)
		return
	}
	response.OkWithMessage("Create users of the group success", c)
}

func DeleteGroupUser(c *gin.Context){
	groupId := c.Param("gid")
	usersId, _ := c.GetQuery("ids")
	err = service.DeleteGroupUser(strings.Split(usersId, ","), groupId)
	if err != nil{
		config.GdoLog.Error("Delete users of the group failed", zap.Any("err", err))
		response.FailWithErrMessage("Delete users of the group failed", err, c)
		return
	}
	response.OkWithMessage("Delete users of the group success", c)
}



func ListPermission(c *gin.Context){
	groups, err := service.ListPermission()
	if err != nil{
		config.GdoLog.Error("List permission failed", zap.Any("err", err))
		response.FailWithErrMessage("List permission failed", err, c)
		return
	}
	response.OkWithDetailed(groups, "List permission success", c)
}

func CreatePermission(c *gin.Context){
	var permission request.Permission
	err = c.ShouldBindJSON(&permission)
	if err != nil {
		response.FailWithErrMessage("Create permission failed", err, c)
		return
	}
	err = service.CreatePermission(permission)
	if err != nil {
		config.GdoLog.Error("Create permission failed", zap.Any("err", err))
		response.FailWithErrMessage("Create permission failed", err, c)
		return
	}
	response.OkWithMessage("Create permission success", c)
}

func ListGroupPermission(c *gin.Context){
	groupId := c.Param("gid")
	permissions, err := service.ListGroupPermission(groupId)
	if err != nil{
		config.GdoLog.Error("List permissions of the group failed", zap.Any("err", err))
		response.FailWithErrMessage("List permissions of the group failed", err, c)
		return
	}
	response.OkWithDetailed(permissions, "List permissions of the group success", c)
}

func CreateGroupPermission(c *gin.Context){
	permissionId := c.Param("pid")
	groupId := c.Param("gid")
	err = service.CreateGroupPermission(permissionId, groupId)
	if err != nil{
		config.GdoLog.Error("Create permissions of the group failed", zap.Any("err", err))
		response.FailWithErrMessage("Create permissions of the group failed", err, c)
		return
	}
	response.OkWithMessage("Create permissions of the group success", c)
}

func DeleteGroupPermission(c *gin.Context){
	groupId := c.Param("gid")
	permissionsId, _ := c.GetQuery("ids")
	err = service.DeleteGroupPermission(strings.Split(permissionsId, ","), groupId)
	if err != nil{
		config.GdoLog.Error("Delete permissions of the group failed", zap.Any("err", err))
		response.FailWithErrMessage("Delete permissions of the group failed", err, c)
		return
	}
	response.OkWithMessage("Delete permissions of the group success", c)
}