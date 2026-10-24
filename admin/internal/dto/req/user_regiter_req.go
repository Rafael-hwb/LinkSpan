package req

type UserRegisterReq struct {
	Username string `json:"username" binding:"required,min=3,max=32"`
	Password string `json:"password" binding:"required,min=8,max=72"`
	RealName string `json:"realName" binding:"required,max=256"`
	Phone    string `json:"phone" binding:"required,len=11,numeric"`
	Mail     string `json:"mail" binding:"required,email,max=512"`
}
