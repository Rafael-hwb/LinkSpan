package toolkit

import (
	"crypto/rand"
	"math/big"
)

const groupGIDCharacters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GenerateGroupGID() (string, error) {
	const length = 6

	var value [length]byte
	limit := big.NewInt(int64(len(groupGIDCharacters)))

	for i := range value {
		index, err := rand.Int(rand.Reader, limit)
		if err != nil {
			return "", err
		}

		value[i] = groupGIDCharacters[index.Int64()]
	}

	return string(value[:]), nil
}
