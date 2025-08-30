package encrypt

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"strings"
)

const (
	saltLength = 16      // Salt length in bytes
	iterations = 100_000 // Number of iterations (recommend >=100,000)
	keyLength  = 32      // Length of derived key in bytes
)

// GeneratePasswordHash generates a PBKDF2 hash with a random salt
func GeneratePasswordHash(password string) (string, error) {
	// Generate random salt
	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// Derive key
	hash := pbkdf2.Key([]byte(password), salt, iterations, keyLength, sha256.New)

	// Encode salt and hash to base64
	saltB64 := base64.RawStdEncoding.EncodeToString(salt)
	hashB64 := base64.RawStdEncoding.EncodeToString(hash)

	// Return in format salt$hash
	return fmt.Sprintf("%s$%s", saltB64, hashB64), nil
}

// ValidatePasswordHash validates a password against a stored salt$hash string
func ValidatePasswordHash(password string, encodedHash string) bool {
	// Split salt and hash
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 2 {
		return false
	}
	saltB64, hashB64 := parts[0], parts[1]

	// Decode salt and hash
	salt, err := base64.RawStdEncoding.DecodeString(saltB64)
	if err != nil {
		return false
	}
	expectedHash, err := base64.RawStdEncoding.DecodeString(hashB64)
	if err != nil {
		return false
	}

	// Derive key from provided password
	hash := pbkdf2.Key([]byte(password), salt, iterations, keyLength, sha256.New)

	// Constant time comparison
	return subtleCompare(hash, expectedHash)
}

// subtleCompare compares two byte slices in constant time
func subtleCompare(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	var result byte
	for i := 0; i < len(a); i++ {
		result |= a[i] ^ b[i]
	}
	return result == 0
}
