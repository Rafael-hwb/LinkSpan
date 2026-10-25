package bootstrap

import (
	"fmt"
	"linkspan/admin/internal/common/security"
	"os"
)

func OpenFieldCipher() (*security.FieldCipher, error) {
	key := os.Getenv("LINKSPAN_FIELD_KEY")

	if key == "" {
		return nil, fmt.Errorf("请设置 LINKSPAN_FIELD_KEY")
	}

	return security.NewFieldCipher(key)

}
