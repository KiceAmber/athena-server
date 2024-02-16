package auth

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" method:"POST" tags:"auth" sm:"博客管理员登录认证"`
	Passport string `json:"passport"`
	Password string `json:"password"`
}

type LoginRes struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}
