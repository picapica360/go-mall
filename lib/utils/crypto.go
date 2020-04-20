package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// MD5 generate MD5 string.
func MD5(p []byte) string {
	h := md5.New()
	h.Write(p)
	cipher := h.Sum(nil)
	return hex.EncodeToString(cipher)
}

// MD5String generate MD5 string.
func MD5String(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	cipher := h.Sum(nil)
	return hex.EncodeToString(cipher)
}
