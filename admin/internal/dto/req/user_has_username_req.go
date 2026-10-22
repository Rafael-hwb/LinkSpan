package req

type UserHasUsernameReq struct {
	Username string `form:"username" binding:"required,max=256"`
}