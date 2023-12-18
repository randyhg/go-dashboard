package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"go-dashboard/Dashboard/controllers"
)

func RegisterRoutes(app *iris.Application) {
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"msg":  "ok",
			"code": 200,
		})
	})
	apiLara(app.Party("api"))
}

func apiLara(router router.Party) {
	router.Post("/sales", controllers.DashboardController.GetSales)
	router.Post("/revenue", controllers.DashboardController.GetRevenue)
	router.Post("/customers", controllers.DashboardController.GetCustomers)
}
