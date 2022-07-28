package cryptography

import (
	"crypto/md5"
	"encoding/hex"
)

type MD5Encryptor struct{}

func NewMD5Encryptor() *MD5Encryptor {
	return &MD5Encryptor{}
}

func (f *MD5Encryptor) Encrypt(value string) string {
	hash := md5.Sum([]byte(value))
	return hex.EncodeToString(hash[:])
}
