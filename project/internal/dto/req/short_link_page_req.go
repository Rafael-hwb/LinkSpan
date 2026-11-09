package req

type ShortLinkPageReq struct {
	GID     string `form:"gid" binding:"required,max=32"`
	Current int32  `form:"current" binding:"required,min=1"`
	Size    int32  `form:"size" binding:"required,min=1,max=100"`
}