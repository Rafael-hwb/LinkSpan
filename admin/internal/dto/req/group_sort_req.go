package req

type GroupSortReq struct {
	GID       string `json:"gid" binding:"required,max=32"`
	SortOrder *int32 `json:"sortOrder" binding:"required"`
}
