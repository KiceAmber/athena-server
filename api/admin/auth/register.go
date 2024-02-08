package auth

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta `path:"/login" method:"POST" tags:"auth" sm:"博客管理员登录认证"`
}

type RegisterRes struct {
}
