package handler

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/token"
	"my_chat/im/ws/internal/svc"
	"my_chat/pkg/ctxdata"
	"net/http"
)

type JwtAuth struct {
	svc    *svc.ServiceContext
	parser *token.TokenParser
	logx.Logger
}

func NewJwtAuth(svc *svc.ServiceContext) *JwtAuth {
	return &JwtAuth{
		svc:    svc,
		parser: token.NewTokenParser(),
		Logger: logx.WithContext(context.Background()),
	}
}

func (j JwtAuth) Auth(w http.ResponseWriter, r *http.Request) bool {
	token, err := j.parser.ParseToken(r, j.svc.Config.JwtAuth.AccessSecret, "")
	if err != nil {
		j.Errorf("parse token err %v", err)
		return false
	}

	if !token.Valid {
		return false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}
	*r = *r.WithContext(context.WithValue(r.Context(), ctxdata.Identify, claims[ctxdata.Identify]))
	return true
}

func (j JwtAuth) UserId(r *http.Request) string {
	return ctxdata.GetUID(r.Context())
}
