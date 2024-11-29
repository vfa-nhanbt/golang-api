package helpers

import "golang.org/x/crypto/bcrypt"

func NormalizePassword(p string) []byte {
	return []byte(p)
}

func EncodeUserPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareEncodePassword(password, encodePassword string) bool {
	passwordByte := []byte(password)
	encodePasswordByte := []byte(encodePassword)
	if err := bcrypt.CompareHashAndPassword(encodePasswordByte, passwordByte); err != nil {
		return false
	}
	return true
}
