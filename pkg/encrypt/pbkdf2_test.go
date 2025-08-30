package encrypt

import (
	"testing"
)

func Test_GenerateAndValidatePasswordHash(t *testing.T) {
	password := "secure-password"

	// Generate hash
	hash, err := GeneratePasswordHash(password)
	if err != nil {
		t.Error(err)
	}
	t.Log("Password hash:", hash)

	// Validate password
	ok := ValidatePasswordHash(password, hash)
	t.Log("Password valid:", ok)

	// Validate wrong password
	ok = ValidatePasswordHash("wrong-password", hash)
	t.Log("Wrong password valid:", ok)
}
