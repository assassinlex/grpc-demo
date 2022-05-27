package token

import (
	"crypto/rsa"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// JwtTokenVerifier jwt token 验证器
type JwtTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

// Verify token 校验 & 返回用户 account_id
func (v *JwtTokenVerifier) Verify(token string) (string, error) {
	// 尝试解析 token
	tokenParsed, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return v.PublicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("无法解析 token: %v", err)
	}
	// token 是否有效
	if !tokenParsed.Valid {
		return "", fmt.Errorf("非法 token: %v", err)
	}
	// 确保 token 的 claims 是解析时传递的 claim 类型: &jwt.StandardClaims{}
	claims, ok := tokenParsed.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("错误的 claim 类型")
	}
	// claim 是否合法: claim 是否被修改、是否已经过期等
	if err := claims.Valid(); err != nil {
		return "", fmt.Errorf("非法的 claim: %v", err)
	}

	return claims.Subject, nil
}
