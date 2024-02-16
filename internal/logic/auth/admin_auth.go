package auth

import (
	"athena-server/internal/consts"
	"athena-server/internal/dao"
	"athena-server/internal/model"
	"athena-server/internal/model/entity"
	"athena-server/internal/service"
	"context"
)

type sAuth struct{}

func init() {
	service.RegisterAuth(New())
}

func New() *sAuth {
	return &sAuth{}
}

func (s sAuth) Login(ctx context.Context, in *model.LoginInput) (out *model.LoginOutput, err error) {
	out = &model.LoginOutput{}

	token := "amber"

	var user = &entity.User{}
	if err = dao.User.Ctx(ctx).
		Where("passport = ? and password = ?", in.Passport, in.Password).
		Scan(&user); err != nil {
		return nil, err
	}

	if user == nil {
		return nil, consts.ErrUserNotFound
	}

	out.Id = user.Id
	out.Token = token
	return
}
