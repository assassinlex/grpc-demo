package auth

import (
	"context"
	"cool_car/shared/auth/token"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"github.com/dgrijalva/jwt-go"

	"google.golang.org/grpc"
)

const (
	bearerPrefix        = "Bearer "
	authorizationHeader = "authorization"
)

var unauthenticated = status.Error(codes.Unauthenticated, "")

// Interceptor 生成拦截器
func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	file, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("open public key file failed: %v", err)
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read public key file failed: %v", err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("parse public key file failed: %v", err)
	}
	i := &interceptor{
		verifier: &token.JwtTokenVerifier{
			PublicKey: key,
		},
	}
	return i.HandleRequest, nil
}

// token 验证器
type tokenVerifier interface {
	Verify(token string) (string, error)
}

// 拦截器
type interceptor struct {
	verifier tokenVerifier
}

// HandleRequest 拦截器职责
func (i *interceptor) HandleRequest(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {

	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return nil, unauthenticated
	}

	// token 换 account_id
	accountID, err := i.verifier.Verify(tkn)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token is invalid: %v", err)
	}

	return handler(ContextWithAccountID(ctx, AccountID(accountID)), req)
}

type accountIDKey struct{}
type AccountID string

func (a AccountID) String() string {
	return string(a)
}

// ContextWithAccountID context 注入 account_id
func ContextWithAccountID(ctx context.Context, aid AccountID) context.Context {
	return context.WithValue(ctx, accountIDKey{}, aid)
}

// AccountIDFromContext 从上线文中获取 account_id
func AccountIDFromContext(ctx context.Context) (AccountID, error) {
	v := ctx.Value(accountIDKey{})
	aid, ok := v.(AccountID)
	if !ok {
		return "", unauthenticated
	}
	return aid, nil
}

// 从外部上下文获取 token
func tokenFromContext(ctx context.Context) (string, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", unauthenticated
	}

	tkn := ""

	for _, v := range md[authorizationHeader] {
		if strings.HasPrefix(v, bearerPrefix) {
			tkn = v[len(bearerPrefix):]
			break
		}
	}

	if tkn == "" {
		return "", unauthenticated
	}

	return tkn, nil
}
