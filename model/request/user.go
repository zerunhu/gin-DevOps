package request

type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
}

type CreateUser struct {
	Username  string `json:"username" binding:"required"   `      // 用户名
	Password  string `json:"password" binding:"required"`     // 密码
	Phone     string `json:"phone" binding:"required,len=11"` // 电话
	Email     string `json:"email" binding:"required,email"`  // 邮箱
}
//type CreateUser struct {
//	User CreateUser
//}
type GroupUser struct {
	Name  string `json:"name" binding:"required"`     // 组名
	Desc  string `json:"desc" binding:"required"`     // 描述
}

