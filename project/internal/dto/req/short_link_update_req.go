package req

type ShortLinkUpdateReq struct {
	OriginGID     string  `json:"originGid" binding:"required,max=32"`
	GID           string  `json:"gid" binding:"required,max=32"`
	FullShortURL  string  `json:"fullShortUrl" binding:"required,max=135"`
	OriginURL     string  `json:"originUrl" binding:"required,max=1024"`
	ValidDateType *int    `json:"validDateType" binding:"required,oneof=0 1"`
	ValidDate     *string `json:"validDate"`
	Describe      string  `json:"describe" binding:"max=1024"`	
}