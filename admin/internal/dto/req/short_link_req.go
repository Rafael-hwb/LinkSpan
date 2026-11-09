package req

type ShortLinkCreateReq struct {
	Domain        string  `json:"domain" binding:"required,max=128"`
	OriginURL     string  `json:"originUrl" binding:"required,max=1024"`
	GID           string  `json:"gid" binding:"required,max=32"`
	CreatedType   *int    `json:"createdType" binding:"required,oneof=0 1"`
	ValidDateType *int    `json:"validDateType" binding:"required,oneof=0 1"`
	ValidDate     *string `json:"validDate"`
	Describe      string  `json:"describe" binding:"max=1024"`
}

type ShortLinkPageReq struct {
	GID     string `form:"gid" binding:"required,max=32"`
	Current int32  `form:"current" binding:"required,min=1"`
	Size    int32  `form:"size" binding:"required,min=1,max=100"`
}