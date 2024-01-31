package middleware

import "github.com/gogf/gf/v2/net/ghttp"

// CorsMiddleware the middleware supported by GoFrame
func CorsMiddleware(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
