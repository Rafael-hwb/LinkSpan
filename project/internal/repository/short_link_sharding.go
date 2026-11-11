package repository

import (
	"fmt"
	"unicode/utf16"
)

const ShortLinkShardCount = 16

func JavaStringHash(value string) int32 {
	var hash int32

	for _, unit := range utf16.Encode([]rune(value)) {
		hash = 31*hash + int32(unit)
	}

	return hash
}

func ShortLinkShardIndex(value string) int64 {
	hash := int64(JavaStringHash(value))
	if hash < 0 {
		hash = -hash
	}

	return hash % int64(ShortLinkShardCount)
}

func ShortLinkTableName(gid string) string {
	return fmt.Sprintf("t_link_%d", ShortLinkShardIndex(gid))
}

func ShortLinkGotoTableName(fullShortURL string) string {
	return fmt.Sprintf("t_link_goto_%d", ShortLinkShardIndex(fullShortURL))
}
