package middleware

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Auth() *gtoken.GfToken {
	gfToken := &gtoken.GfToken{
		LoginPath:       "/login",
		LoginBeforeFunc: Login,
		LogoutPath:      "/logout",
	}

	return gfToken
}

func Login(r *ghttp.Request) (string, interface{}) {
	passport := r.Get("passport")
	password := r.Get("password")

	if passport.String() == "" || password.String() == "" {
		r.Response.WriteJsonExit(gtoken.Fail("账号或密码错误"))
	}
	return passport.String(), "1"
}
