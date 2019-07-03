package main

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

func stringSign(data, key string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	return hex.EncodeToString(hmac.Sum(nil))
}

func signedStringCheck(data, sign, key string) bool {
	return sign == stringSign(data, key)
}