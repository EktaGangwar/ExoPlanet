package controllers

import (
	"encoding/json"
	"exoplanet-service/models"
	"strconv"

	"github.com/astaxie/beego"
)

type ExoplanetController struct {
	beego.Controller
}

func (c *ExoplanetController) Post() {
	log.Println("Request Body:", string(c.Ctx.Input.RequestBody))
	contentType := c.Ctx.Input.Header("Content-Type")
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Content-Type - ", contentType)
	request := c.Ctx.Input.RequestBody
	var exoplanet models.Exoplanet
	if err := json.Unmarshal(request, &exoplanet); err != nil {
		log.Println("Error ", err.Error())
		c.CustomAbort(400, "Invalid JSON")
		return
	}

	exoplanet, err := models.AddExoplanet(exoplanet)
	if err != nil {
		c.CustomAbort(400, err.Error())
		return
	}

	c.Data["json"] = exoplanet
	c.ServeJSON()
}
func (c *ExoplanetController) Get() {
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		exoplanets := make([]models.Exoplanet, 0, len(models.Exoplanets))
		for _, exoplanet := range models.Exoplanets {
			exoplanets = append(exoplanets, exoplanet)
		}
		c.Data["json"] = exoplanets
	} else {
		exoplanet, exists := models.Exoplanets[id]
		if !exists {
			c.CustomAbort(404, "Exoplanet not found")
		}
		c.Data["json"] = exoplanet
	}
	c.ServeJSON()
}

func (c *ExoplanetController) Put() {
	id := c.Ctx.Input.Param(":id")
	var exoplanet models.Exoplanet
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &exoplanet); err != nil {
		c.CustomAbort(400, err.Error())
	}

	exoplanet, err := models.UpdateExoplanet(id, exoplanet)
	if err != nil {
		c.CustomAbort(400, err.Error())
	}

	c.Data["json"] = exoplanet
	c.ServeJSON()
}

func (c *ExoplanetController) Delete() {
	id := c.Ctx.Input.Param(":id")
	if _, exists := models.Exoplanets[id]; !exists {
		c.CustomAbort(404, "Exoplanet not found")
	}
	delete(models.Exoplanets, id)
	c.Data["json"] = map[string]string{"result": "success"}
	c.ServeJSON()
}

func (c *ExoplanetController) GetFuelEstimation() {
	id := c.Ctx.Input.Param(":id")
	exoplanet, exists := models.Exoplanets[id]
	if !exists {
		c.CustomAbort(404, "Exoplanet not found")
	}

	crewCapacityStr := c.GetString("crewCapacity")
	crewCapacity, err := strconv.Atoi(crewCapacityStr)
	if err != nil {
		c.CustomAbort(400, "invalid crew capacity")
	}

	gravity := models.GetGravity(exoplanet)
	fuel := models.CalculateFuel(exoplanet.Distance, gravity, crewCapacity)
	c.Data["json"] = map[string]float64{"fuel_estimation": fuel}
	c.ServeJSON()
}
