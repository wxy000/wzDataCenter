package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"strconv"
	"time"
	"wzDataCenter/common"
	"wzDataCenter/models"
)

type CustomClaims struct {
	models.Users
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte
}

// NewJWT 获取签名 密钥
// var MySecret = []byte("DigiWin-zhejiangJiaofu-Tfuwu")
func NewJWT() *JWT {
	return &JWT{
		[]byte(common.CONF.Jwt.Secret),
	}
}

// GenToken 创建Token
func (j *JWT) GenToken(user models.Users) (string, error) {
	expiresat := common.CONF.Jwt.ExpiresAT
	expiresatint64, err := strconv.ParseInt(expiresat, 10, 64)
	if err != nil {
		return "", err
	}
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(expiresatint64))), //1小时后过期
			Issuer:    common.CONF.Jwt.Issuer,                                            //签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		fmt.Println(" token parse err:", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新Token
func (j *JWT) RefreshToken(tokenStr string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = jwt.At(time.Now().Add(time.Minute * 10))
		return j.GenToken(claims.Users)
	}
	return "", errors.New("couldn't handle this token")
}
