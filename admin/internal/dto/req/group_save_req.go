package req

type GroupSaveReq struct {
	Name string `json:"name" binding:"required,max=64"`
}
