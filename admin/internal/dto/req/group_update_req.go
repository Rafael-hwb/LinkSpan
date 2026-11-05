package req

type GroupUpdateReq struct {
	GID  string `json:"gid" binding:"required,max=32"`
	Name string `json:"name" binding:"required,max=64"`
}