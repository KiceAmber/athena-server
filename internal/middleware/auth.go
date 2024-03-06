package middleware

import (
	"athena-server/internal/consts"
	"athena-server/internal/model"
	"athena-server/internal/service"
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

	// 这里调用 logic 代码
	_, err := service.Auth().Login(r.GetCtx(), &model.LoginInput{
		Passport: passport.String(),
		Password: password.String(),
	})
	if err != nil {
		r.Response.WriteJsonExit(gtoken.Fail(consts.ErrUserNotFound.Error()))
	}

	return passport.String(), "1"
}
