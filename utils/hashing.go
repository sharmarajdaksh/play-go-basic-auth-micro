package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

// GenerateSalt generates a randomized salt of specified length
func GenerateSalt(saltSize int) (string, error) {
	salt := make([]byte, saltSize)

	if _, e := rand.Read(salt[:]); e != nil {
		return "", e
	}

	return base64.URLEncoding.EncodeToString(salt), nil
}

// GenerateSaltedHash generates a hash for data with added salt
func GenerateSaltedHash(data string, salt string) (string, error) {
	saltBytes, e := base64.URLEncoding.DecodeString(salt)
	if e != nil {
		return "", e
	}

	dataBytes := []byte(data)

	dataBytes = append(dataBytes, saltBytes...)

	sha512Hasher := sha512.New()
	if _, e := sha512Hasher.Write(dataBytes); e != nil {
		return "", e
	}

	dataHash := sha512Hasher.Sum(nil)

	b64EncodedDataHash := base64.URLEncoding.EncodeToString(dataHash)

	return b64EncodedDataHash, nil
}

// VerifyHash verifies a given hash with the hash of a string given salt
func VerifyHash(data string, salt string, hash string) (bool, error) {
	generatedHash, e := GenerateSaltedHash(data, salt)
	if e != nil {
		return false, e
	}

	return hash == generatedHash, nil

}

// func DOIT() {

// 	p := "HELLOWORLD"
// 	s, _ := generateSalt(config.C.Security.SaltSize)

// 	h, _ := generateHash(p, s)

// 	v, _ := VerifyHash(p, base64.URLEncoding.EncodeToString(s), h)

// 	fmt.Println(v)
// }
