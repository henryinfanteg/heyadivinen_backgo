package random

import (
	"crypto/rand"
)

// Variables globales
const (
	TypeNumbers           = 1
	TypeLetters           = 2
	TypeLowercaseLetters  = 3
	TypeUppercaseLetters  = 4
	TypeLettersAndNumbers = 5
	TypeAll               = 6
)

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(typeCharacters int, numberCharacters int) (string, error) {
	characters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+-*/@=?()&%$#.,"

	switch typeCharacters {
	case TypeNumbers:
		characters = "0123456789"
	case TypeLetters:
		characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case TypeLowercaseLetters:
		characters = "abcdefghijklmnopqrstuvwxyz"
	case TypeUppercaseLetters:
		characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case TypeLettersAndNumbers:
		characters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	bytes, err := generateRandomBytes(numberCharacters)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = characters[b%byte(len(characters))]
	}
	return string(bytes), nil
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
