package admin

import (
	"athena-server/api/admin/user"
	"context"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (u *User) GetUserList(ctx context.Context, req *user.AddUserReq) (res *user.AddUserRes, err error) {

	return
}
