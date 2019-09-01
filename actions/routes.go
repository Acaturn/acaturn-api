package actions

import (
	"acaturn_api/actions/Api/V1"
	"github.com/gobuffalo/buffalo"
)

func mountRoutes(app *buffalo.App) {
	var api = app.Group("/api")
	var v1 = api.Group("/v1")
	v1.GET("/", V1.HomeHandler)
}
