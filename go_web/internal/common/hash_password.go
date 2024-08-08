package common

import (
	"crypto/md5"
	"encoding/hex"
)

func HashPassword(password string) string {
	bytes := []byte(password)
	hash := md5.Sum(bytes)
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func CheckHashPassword(password, hash string) bool {
	passwordHash := HashPassword(password)
	return passwordHash == hash
}
