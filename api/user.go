package api

import (
	"gin-DevOps/config"
	"gin-DevOps/middleware"
	"gin-DevOps/model"
	"gin-DevOps/model/request"
	"gin-DevOps/model/response"
	"gin-DevOps/service"
	"gin-DevOps/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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
	// 用户发送用户名和密码过来
	var user request.Login
	err := c.ShouldBindJSON(&user)
	if err != nil {
		config.GdoLog.Error("登录失败",zap.Any("err",err))
		response.FailWithMessage("登录失败,"+err.Error(), c)
		return
	}
	// 校验用户名和密码是否正确
	sqlUser := model.User{}
	config.GdoDb.Where("username = ?", user.Username).Find(&sqlUser)
	password := utils.MD5V([]byte(user.Password))
	if password == sqlUser.Password {
		// 生成Token
		tokenString, _ := middleware.GenToken(user.Username)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"token": tokenString,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 500,
		"msg":  "账号或者密码错误",
	})
	return
}

func UserInfo(c *gin.Context) {
	username := c.MustGet("username").(string)
	roles := []string{"GROUP_PERMISSION_UPDATE","USERS_LIST","GROUPS_UPDATE","OPERATIONPLAN_CREATE","PROD_WORLD_LIST","PERMISSION_LIST","QA_WORLD_DELETE","PROD_WORLD_HISTORY","OPERATIONPLAN_RETRIEVE","PROD_WORLD_DELETE","PROD_WORLD_CLIENTLIST","PROD_WORLD_SERVERLISTONLINE","CIPROCESS_CREATE","OPERATIONPLAN_UPDATE","DEV_WORLD_HISTORY","QA_WORLD_LIST","QA_WORLD_STATUS","PROD_WORLD_CLEARDEADNUMBER","PROD_WORLD_CREATE","USERS_DELETE","PROD_WORLD_RESTART","QA_WORLD_HISTORY","DEV_NODEGROUP_LIST","CI_IMAGE_LIST","DEV_WORLD_STATUS","USERS_CREATE","GROUP_USER_DELETE","CIPROCESS_DELETE","QA_WORLD_RESTART","DEV_WORLD_RESTART","PROD_WORLD_UPDATE","PROD_WORLD_STATUS","QA_WORLD_CREATE","GROUPS_LIST","DEV_WORLD_BACKUP","DEV_NODEGROUP_UPDATE","PROD_WORLD_SERVERLIST","CIPROCESS_PUSH","PROD_WORLD_NOTICE","CIPROCESS_LIST","QA_WORLD_UPDATE","PROD_WORLD_BACKUP","DEV_WORLD_CREATE","QA_WORLD_BACKUP","OPERATIONPLAN_COMPLETE","GROUP_PERMISSION_LIST","PROD_WORLD_SECURITYGROUP","PROD_WORLD_AUTOCLEARDEADNUMBER","GROUP_USER_LIST","SVNINFO_LIST","GROUPS_DELETE","PERMISSION_CREATE","PERMISSION_DELETE","OPERATIONPLAN_LIST","DEV_WORLD_UPDATE","OPERATIONPLAN_DELETE","PROD_NODEGROUP_UPDATE","GROUP_USER_CREATE","PROD_NODEGROUP_LIST","DEV_WORLD_DELETE","GROUPS_CREATE","DEV_WORLD_LIST","CIPROCESS_BUILD"}
	c.JSON(http.StatusOK, gin.H{
		"user_name": username,
		"roles": roles,
	})
}

func ListUser(c *gin.Context) {
	var users []model.User
	if err := config.GdoDb.Find(&users).Error; err != nil {
		config.GdoLog.Error("获取用户列表失败",zap.Any("err",err))
		response.FailWithMessage("获取用户列表失败,"+err.Error(), c)
		return
	}
	group := "dev"
	var responseUser []response.UserListResponse
	for _, user := range users{
		u := response.UserListResponse{
			Id:    user.ID,
			Name:  user.Username,
			Email: user.Email,
			Phone: user.Phone,
			Group: group,
		}
		responseUser = append(responseUser, u)
	}
	response.OkWithDetailed(responseUser, "获取用户列表成功", c)
}

func CreateUser(c *gin.Context) {
	var U request.CreateUser
	err := c.ShouldBindJSON(&U)
	if err != nil {
		config.GdoLog.Error("注册失败",zap.Any("err",err))
		response.FailWithMessage("注册失败,"+err.Error(), c)
		return
	}
	user := &model.User{Username: U.User.Username, Password: U.User.Password, Email: U.User.Email, Phone: U.User.Phone}
	err, userReturn := service.Register(*user)
	if err != nil {
		config.GdoLog.Error("注册失败", zap.Any("err", err))
		response.FailWithMessage("注册失败,"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.RegisterUserResponse{User: userReturn}, "注册成功", c)
	}
}

func DeleteUser(c *gin.Context){
	var user model.User
	userId := c.Param("id")
	if err := config.GdoDb.Where("id = ?", userId).Delete(&user).Error; err != nil {
		config.GdoLog.Error("删除用户失败",zap.Any("err",err))
		response.FailWithMessage("删除用户失败,"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}