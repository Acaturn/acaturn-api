package actions

import (
	"acaturn_api/actions/api/version1"
	"github.com/gobuffalo/buffalo"
)

func mountRoutes(app *buffalo.App) {
	var api = app.Group("/api")
	var v1 = api.Group("/version1")
	v1.GET("/", version1.HomeHandler)
}
