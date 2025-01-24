package test

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestGenerateFromPassword(t *testing.T) {
	password, err := bcrypt.GenerateFromPassword([]byte("1612977482"), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	fmt.Println(string(password))
}

func TestCompareHashAndPassword(t *testing.T) {
	//password := "$2a$10$iqLH2H1h/.hyfZA3PSrYjuyC/vi6.O9TtsMHpsIchz8DVQmTHcE4e"
	// $2a$10$pALp0RD/pKL7.3Zt8gHfZ.oDb5WG1v/q7KYodPbEWCqOyxxjlQQOW
	// $2a$10$Mf5U5cMS/c3KFYG4gSjTf.iwycC00ari2zOXwnHABNiZSsQs8mTTO
	// $2a$10$ZVD/Qcc2JcWCf1TAKJ9dhezthk4NYxsMEgd7grtU1M.e8uCNdaTwC
	// $2a$10$KglYb61RmW1L3GW8FPVl1O592H6z7Ihx2Ex8MocXiMgPJ1pojzLwC
	password := "$2a$10$bMQ/gquBUtBOREBLZ0upNeeSfviv7bEZD.hdPCPNB9a0EKh6EkAG2"
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte("1612977482"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("密码正确")
}
