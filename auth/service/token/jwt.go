package token

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtTokenGenerator struct {
	// 发布者
	issuer string
	// 当前时间函数
	nowFunc func() time.Time
	// 私钥
	privateKey *rsa.PrivateKey
}

// NewJwtTokenGenerator 构造函数
func NewJwtTokenGenerator(issuer string, key *rsa.PrivateKey) *JwtTokenGenerator {
	return &JwtTokenGenerator{
		issuer:     issuer,
		nowFunc:    time.Now,
		privateKey: key,
	}
}

func (t *JwtTokenGenerator) GenerateToken(accountID string, expire time.Duration) (string, error) {
	// 这里将获取当前时间作为方法封装起来, 由外部注入, 方便写表格驱动测试
	// nowSec := time.Now().Unix()
	nowSec := t.nowFunc().Unix()
	// 获取未签名的 token
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		ExpiresAt: nowSec + int64(expire.Seconds()),
		IssuedAt:  nowSec,
		Issuer:    t.issuer,
		Subject:   accountID,
	})
	// 用私钥对 token 进行签名
	// 返回签名
	return token.SignedString(t.privateKey)
}
