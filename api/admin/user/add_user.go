package user

import "github.com/gogf/gf/v2/frame/g"

type AddUserReq struct {
	g.Meta   `path:"/addUser" method:"POST" tags:"user" sm:"添加用户"`
	Passport string `json:"passport"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AddUserRes struct {
}
