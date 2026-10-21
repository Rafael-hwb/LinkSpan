package resp

//查询用户接口相应体
type UserResp struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	RealName string `json:"realName"`
	Phone    string `json:"phone"`
	Mail     string `json:"mail"`
}
