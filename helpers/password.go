package helpers

import "golang.org/x/crypto/bcrypt"

// Bcrypt 密码加密
func BcryptPasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Bcrypt 密码验证
func BcryptPasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
