package middleware

import "github.com/gogf/gf/v2/net/ghttp"

func Auth(r *ghttp.Request) {

}

func Cors(r *ghttp.Request) {
	CorsMiddleware(r)
}
