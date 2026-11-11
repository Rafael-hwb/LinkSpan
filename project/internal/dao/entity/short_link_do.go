package entity

import "time"

type ShortLinkDO struct {
	Domain        string
	ShortURI      string
	FullShortURL  string
	OriginURL     string
	GID           string
	Favicon       string
	CreatedType   int
	ValidDateType int
	ValidDate     *time.Time
	Describe      string
}
