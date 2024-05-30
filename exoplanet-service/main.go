package main

import (
	_ "exoplanet-service/routers"

	"github.com/astaxie/beego"
)

var RedirectHttp = func(ctx *context.Context) {
	if !ctx.Input.IsSecure() {
		// no need for an additional '/' between domain and uri
		url := "https://" + ctx.Input.Domain() + ":" + beego.AppConfig.String("HttpsPort") + ctx.Input.URI()
		ctx.Redirect(302, url)
	}
}

func main() {
	beego.Run()
}
