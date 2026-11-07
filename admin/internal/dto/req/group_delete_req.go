package req

type GroupDeleteReq struct {
	GID string `form:"gid" binding:"required,max=32"`
}
