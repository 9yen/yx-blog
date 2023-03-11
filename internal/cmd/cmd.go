package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"yx-blog/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			//err = s.SetConfigWithMap(g.Map{
			//	"address":          ":8000",
			//	"indexFiles":       g.Slice{"index.html", "main.html"},
			//	"accessLogEnabled": true,
			//	"errorLogEnabled":  true,
			//	"pprofEnabled":     true,
			//	"sessionIdName":    "MySessionId",
			//	"sessionMaxAge":    24 * time.Hour,
			//	"dumpRouterMap":    false,
			//})
			//if err != nil {
			//	return err
			//}
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,
					controller.User,
				)
			})
			s.Run()
			return
		},
	}
)
