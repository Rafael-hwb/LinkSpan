package repository

import (
	"fmt"
	"unicode/utf16"
)

//当前固定的用户物理表数量
const UserShardCount = 16

//Java String.hashCode 规则计算哈希
func javaStringHash(value string) int32 {
	var hash int32

	for _, unit := range utf16.Encode([]rune(value)) {
		hash = 31*hash + int32(unit)
	}

	return hash
}

//返回用户名对应的物理表名
func UserTableName(username string) string {
	hash := javaStringHash(username)

	value := int64(hash)
	if value < 0 {
		value = -value
	}

	index := value % int64(UserShardCount)

	return fmt.Sprintf("t_user_%d", index)
}