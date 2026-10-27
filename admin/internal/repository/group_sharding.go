package repository

import "fmt"

const GroupShardCount = 16

func GroupTableName(username string) string {
	value := int64(javaStringHash(username))

	if value < 0 {
		value = -value
	}

	index := value % int64(GroupShardCount)

	return fmt.Sprintf("t_group_%d", index)
}