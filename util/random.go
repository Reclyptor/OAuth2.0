package util

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"strings"
)

func GenerateBytes(n int) []byte {
	bytes := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, bytes)
	if err != nil {
		panic(err)
	}
	return bytes
}

func GenerateHex(length int) string {
	bytes := GenerateBytes(length / 2)
	return strings.ToUpper(hex.EncodeToString(bytes))
}
