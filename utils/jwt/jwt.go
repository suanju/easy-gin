package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Hotkey 密钥
var Hotkey = []byte("G0-store")

//Claims  TOKEN 的结构体
type Claims struct {
	Uid uint
	jwt.StandardClaims
}

// NextToken 登录以后签发jwt
func NextToken(uid uint) string {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			//过期时间
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			// 签名颁发者
			Issuer: "root",
			//签名主题
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(Hotkey)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

// ParseToken 解析 Token
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Hotkey, nil
	})
	if err != nil {
		fmt.Println(" token parse err:", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
