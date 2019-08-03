package main

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/swaggo/gogf-swagger"
	"github.com/swaggo/gogf-swagger/swaggerFiles"

	_ "github.com/swaggo/gogf-swagger/example/docs"

)

//
// @Summary hello接口
// @Description hello接口
// @Tags hello
// @Success 200 {string} string	"ok"
// @router / [get]
// @Security ApiKeyAuth
func hello(r *ghttp.Request) {
	r.Response.Writeln("哈喽世界！")
}


func run() {
	s := g.Server()

	s.BindHookHandlerByMap("/*", map[string]ghttp.HandlerFunc{
		"BeforeServe": func(r *ghttp.Request) {
			r.Response.CORS(ghttp.CORSOptions{
				AllowOrigin:      "*",
				AllowMethods:     ghttp.HTTP_METHODS,
				AllowCredentials: "true",
				MaxAge:           3628800,
				AllowHeaders:"*",
			})
		},
	})
	s.BindHandler("/", hello)

	url := gogfSwagger.URL("http://localhost:8199/swagger/doc.json") //The url pointing to API definition
	s.BindHandler("/swagger/*any", gogfSwagger.WrapHandler(swaggerFiles.Handler, url))

	s.SetAccessLogEnabled(true)
	s.SetPort(8199)
	s.Run()
}