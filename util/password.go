package util

const SALT_LENGTH = 16
const ITERATIONS = 100000

func HashPassword(password string) string {
	salt := GenerateBytes(SALT_LENGTH)
	hash := []byte(password)
	for i := 0; i < ITERATIONS-1; i++ {
		hash = SHA256(hash, salt)[SALT_LENGTH:]
	}
	return SHA256Hex(hash, salt)
}

func ValidatePassword(password string, hex string) bool {
	salt, _ := SHA256Unhex(hex)
	hash := []byte(password)
	for i := 0; i < ITERATIONS-1; i++ {
		hash = SHA256(hash, salt)[SALT_LENGTH:]
	}
	return hex == SHA256Hex(hash, salt)
}
