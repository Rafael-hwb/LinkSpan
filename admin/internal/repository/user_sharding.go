package repository

import (
	"fmt"
	"unicode/utf16"
)

// 哈希定位用户表
func UserTableName(username string) string {
	var hash int32

	//utf-16编码单元计算
	for _, unit := range utf16.Encode([]rune(username)) {
		hash = 31*hash + int32(unit)
	}

	value := int64(hash)
	if value < 0 {
		value = -value
	}

	return fmt.Sprintf("t_user_%d", value%16)
}
