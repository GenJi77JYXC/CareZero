package utils

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func Encrypt(encryptText string) (string, error) {
	hashStr, err := bcrypt.GenerateFromPassword([]byte(encryptText), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashStr), err
}

func CompareHashAndPassword(hashPassword, Password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(Password))

	return err == nil
}

func VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 检验手机号
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
