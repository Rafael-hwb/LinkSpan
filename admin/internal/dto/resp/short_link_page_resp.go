package resp

type ShortLinkPageResp struct {
	Records []ShortLinkPageItem `json:"records"`
	Total   int64               `json:"total"`
	Current int32               `json:"current"`
	Size    int32               `json:"size"`	
}

type ShortLinkPageItem struct {
	GID          string `json:"gid"`
	FullShortURL string `json:"fullShortUrl"`
	OriginURL    string `json:"originUrl"`
	Describe     string `json:"describe"`
}