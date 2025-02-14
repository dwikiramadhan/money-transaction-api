package utils

import "golang.org/x/crypto/bcrypt"

// NormalizePassword func for a returning the users input as a byte slice.
func NormalizePassword(p string) []byte {
	return []byte(p)
}

// ComparePasswords func for a comparing password.
func ComparePIN(hashedPin, inputPin string) bool {
	// Since we'll be getting the hashed password from the DB it will be a string,
	// so we'll need to convert it to a byte slice.
	byteHash := NormalizePassword(hashedPin)
	byteInput := NormalizePassword(inputPin)

	// Return result.
	if err := bcrypt.CompareHashAndPassword(byteHash, byteInput); err != nil {
		return false
	}

	return true
}
