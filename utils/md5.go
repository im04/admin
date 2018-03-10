package utils

import (
	"fmt"
	"crypto/md5"
)

func EncryptPassword(password string) (hash string) {
	hash = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	//hash = password
	//hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	return
}