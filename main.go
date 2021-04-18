package main

import (
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "qlsv/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter,
		cors.Allow(&cors.Options{
			AllowAllOrigins: true,
			AllowMethods:    []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
			AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin",
				"content-type", "Content-Type", "sessionkey", ""},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Type"},
			AllowCredentials: true,
		}))
	beego.Run()
}
