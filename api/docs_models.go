package api


type _ResponsePostList struct {
	Code    int                     `json:"code"`    // 业务响应状态码
	Message string                  `json:"message"` // 提示信息
	Data    string                  `json:"data"`    // token
}