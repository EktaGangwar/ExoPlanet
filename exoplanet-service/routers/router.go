package routers

import (
	"exoplanet-service/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/exoplanets", &controllers.ExoplanetController{}, "post:Post")
	beego.Router("/exoplanets", &controllers.ExoplanetController{}, "get:Get")
	beego.Router("/exoplanets/:id", &controllers.ExoplanetController{}, "get:Get")
	beego.Router("/exoplanets/:id", &controllers.ExoplanetController{}, "put:Put")
	beego.Router("/exoplanets/:id", &controllers.ExoplanetController{}, "delete:Delete")
	beego.Router("/exoplanets/:id/fuel", &controllers.ExoplanetController{}, "get:GetFuelEstimation")
}
