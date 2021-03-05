package security

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"io"
)

// EncryptMD5 encripta un texto utilizando el algoritmo MD5
func EncryptMD5(text string) (string, error) {
	md5 := md5.New()
	_, err := io.WriteString(md5, text)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(md5.Sum(nil)), nil
}

// EncryptSHA256 encripta un texto utilizando el algoritmo SHA256
func EncryptSHA256(text string) (string, error) {
	sha256 := sha256.New()
	_, err := sha256.Write([]byte(text))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sha256.Sum(nil)), nil
}

// EncryptSHA512 encripta un texto utilizando el algoritmo SHA512
func EncryptSHA512(text string) (string, error) {
	sha512 := sha512.New()
	_, err := sha512.Write([]byte(text))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sha512.Sum(nil)), nil
}

// EncryptHMACSHA512 encripta un texto utilizando el algoritmo HMAC-SHA512
func EncryptHMACSHA512(text string, key string) string {
	hmac512 := hmac.New(sha512.New, []byte(key))
	hmac512.Write([]byte(text))
	return base64.StdEncoding.EncodeToString(hmac512.Sum(nil))
}
