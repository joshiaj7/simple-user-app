package util

import (
	b64 "encoding/base64"
)

// Encrypt function used to encrypt password
func Encrypt(data []byte) string {
	return b64.StdEncoding.EncodeToString([]byte(data))
}

// Decrypt function used to decrypt password
func Decrypt(data string) string {
	sDec, err := b64.StdEncoding.DecodeString(data)
	if err != nil {
		return "Failed to decode string, error: " + err.Error()
	}

	return string(sDec)
}
