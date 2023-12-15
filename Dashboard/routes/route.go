package routes

import (
	"github.com/kataras/iris/v12"
	"go-dashboard/Dashboard/controllers"
)

func RegisterRoutes(app *iris.Application) {
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"msg":  "ok",
			"code": 200,
		})
	})
	app.Get("/sales", controllers.DashboardController.GetSales)
	app.Get("/revenues", controllers.DashboardController.GetRevenue)
	app.Get("/customers", controllers.DashboardController.GetCustomers)
	app.Get("/recent_sales", controllers.DashboardController.GetRecentSales)
}
