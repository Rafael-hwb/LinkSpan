package req

type UserHasUsernameReq struct {
	usrname string `form:"username" binding:"required,max=256"`
}