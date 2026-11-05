package resp

type GroupResp struct {
	GID            string `json:"gid"`
	Name           string `json:"name"`
	SortOrder      int    `json:"sortOrder"`
	ShortLinkCount *int   `json:"shortLinkCount"`
}
