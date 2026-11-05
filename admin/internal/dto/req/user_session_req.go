package req

// 用于检查登陆状态与退出状态
type UserSessionReq struct {
	Username string `form:"username" binding:"required,max=256"`
	Token    string `form:"token" binding:"required,max=256"`
}
