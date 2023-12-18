package controllers

import (
	"github.com/kataras/iris/v12"
	"go-dashboard/Dashboard/services"
	"go-dashboard/model"
	"math"
	"strconv"
)

type dashboardcontroller struct{}

var DashboardController = new(dashboardcontroller)

func (d *dashboardcontroller) GetSales(ctx iris.Context) {
	var strFilter model.StrFilter
	if err := ctx.ReadForm(&strFilter); err != nil {
		model.FailWithDetailed(err, "Failed to parse JSON Body", ctx)
		return
	}

	filterName, start_date, end_date, start_dateAgo, end_dateAgo, err := services.SwitchCase(strFilter.Filter)
	if err != nil {
		model.FailWithMessage(err.Error(), ctx)
		return
	}

	sales, err := services.DashboardService.GetSalesService(start_date, end_date)
	if err != nil {
		model.FailWithDetailed(err, "Failed get Sales data", ctx)
		return
	}

	percentage, increaseOrDecrease := services.DashboardService.GetSalesIncreaseDecrease(start_date, end_date, start_dateAgo, end_dateAgo)
	if math.IsNaN(percentage) {
		percentage = float64(0)
	}

	data := iris.Map{
		"value":              sales,
		"increasePercentage": strconv.FormatFloat(percentage, 'f', 2, 64),
		"increaseType":       increaseOrDecrease,
		"filter":             filterName,
	}

	model.OkWithData(data, ctx)
}

func (d *dashboardcontroller) GetRevenue(ctx iris.Context) {
	var strFilter model.StrFilter
	if err := ctx.ReadForm(&strFilter); err != nil {
		model.FailWithDetailed(err, "Failed to parse JSON Body", ctx)
		return
	}

	filterName, start_date, end_date, start_dateAgo, end_dateAgo, err := services.SwitchCase(strFilter.Filter)
	if err != nil {
		model.FailWithMessage(err.Error(), ctx)
		return
	}

	revenues, err := services.DashboardService.GetRevenuesService(start_date, end_date)
	if err != nil {
		model.FailWithDetailed(err, "Failed get Revenues data", ctx)
		return
	}
	revenues, err = services.FormatCurrency(revenues)
	if err != nil {
		model.FailWithDetailed(err, "Failed to format Revenue to Rupiah", ctx)
		return
	}

	percentage, increaseOrDecrease := services.DashboardService.GetRevenueIncreaseDecrease(start_date, end_date, start_dateAgo, end_dateAgo)
	if math.IsNaN(percentage) {
		percentage = float64(0)
	}

	data := iris.Map{
		"value":              revenues,
		"increasePercentage": strconv.FormatFloat(percentage, 'f', 2, 64),
		"increaseType":       increaseOrDecrease,
		"filter":             filterName,
	}

	model.OkWithData(data, ctx)
}

func (d *dashboardcontroller) GetCustomers(ctx iris.Context) {
	var strFilter model.StrFilter
	if err := ctx.ReadForm(&strFilter); err != nil {
		model.FailWithDetailed(err, "Failed to parse JSON Body", ctx)
		return
	}

	filterName, start_date, end_date, start_dateAgo, end_dateAgo, err := services.SwitchCase(strFilter.Filter)
	if err != nil {
		model.FailWithMessage(err.Error(), ctx)
		return
	}

	customers, err := services.DashboardService.GetCustomersService(start_date, end_date)
	if err != nil {
		model.FailWithDetailed(err, "Failed get Revenues data", ctx)
		return
	}

	percentage, increaseOrDecrease := services.DashboardService.GetCustomerIncreaseDecrease(start_date, end_date, start_dateAgo, end_dateAgo)
	if math.IsNaN(percentage) {
		percentage = float64(0)
	}

	data := iris.Map{
		"value":              customers,
		"increasePercentage": strconv.FormatFloat(percentage, 'f', 2, 64),
		"increaseType":       increaseOrDecrease,
		"filter":             filterName,
	}

	model.OkWithData(data, ctx)
}
