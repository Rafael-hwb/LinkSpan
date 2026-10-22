package serialize

import "encoding/json"

type Phone string

//手机号json编码
func (p Phone) MarshalJSON() ([]byte, error) {
	masked := maskPhone(string(p))
	return json.Marshal(masked)
}

//隐藏手机号4-7位
func maskPhone(phone string) string {
	chars := []rune(phone)
	for i := 3; i < 7 && i < len(chars); i++ {
		chars[i] = '*'
	}
	return string(chars)
}
