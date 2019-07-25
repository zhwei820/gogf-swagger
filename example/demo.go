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
func hello(r *ghttp.Request) {
	r.Response.Writeln("哈喽世界！")
}


// @title Swagger Example API
// @version 1.0
// @description This is a hello server .
// @termsOfService http://hello.io/terms/

// @contact.name hello
// @contact.url http://www.hello.io
// @contact.email hello@hello.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath /
func main() {
	s := g.Server()
	s.BindHandler("/", hello)

	url := gogfSwagger.URL("http://localhost:8199/swagger/doc.json") //The url pointing to API definition
	s.BindHandler("/swagger/*any", gogfSwagger.WrapHandler(swaggerFiles.Handler, url))

	s.SetPort(8199)
	s.Run()
}