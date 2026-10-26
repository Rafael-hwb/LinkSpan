package req

type UserUpdateReq struct {
	Username string  `json:"username" binding:"required"`
	RealName *string `json:"realName" binding:"omitempty,min=1,max=256"`
	Phone    *string `json:"phone" binding:"omitempty,len=11,numeric"`
	Mail     *string `json:"mail" binding:"omitempty,email,max=512"`
}
