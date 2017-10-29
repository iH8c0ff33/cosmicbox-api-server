package main

import (
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/controllers"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/models"
	"github.com/gin-gonic/gin"
)

const basePath = "/api/v1"

func main() {
	app := gin.Default()
	models.Initialize()

	v1 := app.Group(basePath)
	{
		controllers.Events(v1.Group("/events"))
	}

	app.Run(":9001")
}
