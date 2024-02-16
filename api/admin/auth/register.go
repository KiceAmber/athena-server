package auth

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta   `path:"/register" method:"POST" tags:"auth" sm:"注册管理员身份"`
	Passport string `json:"passport"`
	Password string `json:"password"`
}

type RegisterRes struct {
	Id int `json:"id"`
}
