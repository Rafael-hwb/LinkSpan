package security

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strings"
)

const encryptedPrefix = "enc:vi:"

type FieldCipher struct {
	aead cipher.AEAD
}

func NewFieldCipher(encodeKey string) (*FieldCipher, error) {
	key, err := base64.StdEncoding.DecodeString(encodeKey)
	if err != nil {
		return nil, fmt.Errorf("字段密钥不是有效的 Base64")
	}

	if len(key) != 32 {
		return nil, fmt.Errorf("字段密钥必须是32字节")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCMWithRandomNonce(block)
	if err != nil {
		return nil, err
	}

	return &FieldCipher{
		aead: aead,
	}, nil
}

func fieldAAD(field string) []byte {
	return []byte("linkspan:user:" + field + ":v1")
}

//将正常字段值转换为带版本前缀的密文
func (f *FieldCipher) Encrypt(field string, value string) (string, error) {
	if value == "" {
		return "", nil
	}

	encrypted := f.aead.Seal(
		nil,
		nil,
		[]byte(value),
		fieldAAD(field),
	)

	return encryptedPrefix +
		base64.StdEncoding.EncodeToString(encrypted), nil
}

// Decrypt 验证并解密字段
func (f *FieldCipher) Decrypt(field string, value string) (string, error) {
	if value == "" {
		return "", nil
	}

	if !strings.HasPrefix(value, encryptedPrefix) {
		return "", fmt.Errorf("字段 %s 尚未迁移或密文格式不支持", field)
	}

	encoded := strings.TrimPrefix(value, encryptedPrefix)

	encrypted, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", fmt.Errorf("字段 %s 密文编码无效", field)
	}

	plaintext, err := f.aead.Open(
		nil,
		nil,
		encrypted,
		fieldAAD(field),
	)
	if err != nil {
		return "", fmt.Errorf("字段 %s 密文验证失败: %w", field, err)
	}

	return string(plaintext), nil
}
