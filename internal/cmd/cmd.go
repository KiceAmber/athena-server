package cmd

import (
	"athena-server/internal/controller/admin"
	"athena-server/internal/controller/blog"
	"athena-server/internal/middleware"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api", func(apiGroup *ghttp.RouterGroup) {
				apiGroup.Middleware(ghttp.MiddlewareHandlerResponse) // 设置统一响应
				apiGroup.Middleware(middleware.CorsMiddleware)       // Cors 跨域处理

				// 博客后台请求路由组
				apiGroup.Group("/admin", func(adminGroup *ghttp.RouterGroup) {

					// 添加 gToken 中间件
					if err = middleware.Auth().Middleware(ctx, adminGroup); err != nil {
						panic(err)
					}

					adminGroup.Group("/tag", func(group *ghttp.RouterGroup) {
						group.Bind(admin.NewTag())
					})
					adminGroup.Group("/article", func(group *ghttp.RouterGroup) {
						group.Bind(admin.NewArticle())
					})
				})

				// 博客前台请求路由组
				apiGroup.Group("/blog", func(blogGroup *ghttp.RouterGroup) {
					blogGroup.Group("/article", func(group *ghttp.RouterGroup) {
						group.Bind(blog.NewArticle())
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
