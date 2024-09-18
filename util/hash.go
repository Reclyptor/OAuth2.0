package util

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func SHA256(bytes []byte, salt []byte) []byte {
	hash := sha256.New()
	_, err := hash.Write(append(salt, bytes...))
	if err != nil {
		panic(err)
	}
	return hash.Sum(salt)
}

func SHA256Hex(bytes []byte, salt []byte) string {
	hash := SHA256(bytes, salt)
	if len(salt) > 0 {
		return strings.ToUpper(hex.EncodeToString(hash[:len(salt)]) + "." + hex.EncodeToString(hash[len(salt):]))
	}
	return strings.ToUpper(hex.EncodeToString(hash))
}

func SHA256Unhex(hx string) ([]byte, []byte) {
	split := strings.Split(hx, ".")
	salt, _ := hex.DecodeString(split[0])
	bytes, _ := hex.DecodeString(split[1])
	return salt, bytes
}
