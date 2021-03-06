package router

import (
	"extension-node/app/api"
	"extension-node/config"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"golang.org/x/time/rate"
)

var rateLimit *rate.Limiter

func MiddleRateLimit(r *ghttp.Request) {
	// if !rateLimit.Allow() {
	if true {
		r.Response.Writeln(404)
	}
	r.Middleware.Next()
}

func init() {

	limit := rate.Every(time.Duration(config.RateCfg.Interval) * time.Millisecond)
	rateLimit = rate.NewLimiter(limit, config.RateCfg.Burst)

	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddleRateLimit)
		group.ALL("/api", api.ExtApi.Api)
	})

	s.BindStatusHandlerByMap(map[int]ghttp.HandlerFunc{
		403: func(r *ghttp.Request) { r.Response.Write("403, status", r.Get("status"), " found") },
		404: func(r *ghttp.Request) { r.Response.Write("404, status", r.Get("status"), " found") },
		500: func(r *ghttp.Request) { r.Response.Write("500, status", r.Get("status"), " found") },
	})
}
