package resp

type ShortLinkCreateResp struct {
	GID          string `json:"gid"`
	OriginURL    string `json:"originUrl"`
	FullShortURL string `json:"fullShortUrl"`	
}