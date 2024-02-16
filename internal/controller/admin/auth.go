package admin

import (
	"athena-server/api/admin/auth"
	"context"
)

type Auth struct{}

func NewAuth() *Auth {
	return &Auth{}
}

func (a Auth) Login(ctx context.Context, req *auth.LoginReq) (res *auth.LoginRes, err error) {

	return
}
