package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
	"time"
)

const (
	jwtSecret     = "genji11"
	AccessExpire  = time.Hour
	RefreshExpire = time.Hour * 144
)

func GenerateToken(userId uint, role string, rds *redis.Redis) (string, error) {
	refreshToken, err := generateRefreshToken()
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"user_id":     userId,
		"role":        role,
		"refresh":     refreshToken,
		"refresh_exp": time.Now().Add(RefreshExpire).Unix(),
		"access_exp":  time.Now().Add(AccessExpire).Unix(), // 设置过期时间
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	err = storeRefreshTokenInRedis(rds, token, int64(userId), time.Now().Add(RefreshExpire).Unix())
	if err != nil {
		return "", err
	}

	return token, err
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("无效的签名方法")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效的 Token")
}

func generateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func storeRefreshTokenInRedis(rdb *redis.Redis, token string, userID int64, expiresAt int64) error {
	ctx := context.Background()
	err := rdb.SetCtx(ctx, token, strconv.FormatInt(userID, 10))
	if err != nil {
		return err
	}

	err = rdb.ExpireCtx(ctx, token, int(time.Unix(expiresAt, 0).Sub(time.Now()).Seconds()))
	if err != nil {
		return err
	}
	return err
}

func validateRefreshToken(rdb *redis.Redis, token string) (int, error) {
	userID, err := rdb.Get(token)
	if err != nil {
		return 0, err
	}
	id, err := strconv.Atoi(userID)
	if err != nil {
		return 0, err
	}
	return id, nil
}
