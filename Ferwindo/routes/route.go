package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"go-dashboard/Ferwindo/controllers"
)

func RegisterRoutes(app *iris.Application) {
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"msg":  "ok",
			"code": 200,
		})
	})
	apiLara(app.Party("api"))
	apiVendor(app.Party("api/vendor"))
	apiCompany(app.Party("api/company"))
}

func apiLara(router router.Party) {
	router.Post("/sales", controllers.DashboardController.GetSales)
	router.Post("/revenue", controllers.DashboardController.GetRevenue)
	router.Post("/customers", controllers.DashboardController.GetCustomers)

}

func apiVendor(router router.Party) {
	router.Get("/list", controllers.VendorController.GetVendor)
	router.Post("/add", controllers.VendorController.CreateVendor)
	router.Post("/update", controllers.VendorController.UpdateVendor)
	router.Post("/delete", controllers.VendorController.DeleteVendor)
}

func apiCompany(router router.Party) {
	router.Get("/list", controllers.CompanyController.GetCompany)
	router.Post("/add", controllers.CompanyController.CreateCompany)
	router.Post("/update", controllers.CompanyController.UpdateCompany)
	router.Post("/delete", controllers.CompanyController.DeleteCompany)
}
