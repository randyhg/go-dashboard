package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"go-dashboard/Dashboard/services"
	"go-dashboard/model"
	"math"
)

type dashboardcontroller struct{}

var DashboardController = new(dashboardcontroller)
var Today = "today"
var ThisMonth = "this_month"
var ThisYear = "this_year"

func (d *dashboardcontroller) GetSales(ctx iris.Context) {
	var strFilter model.StrFilter
	var start_date string
	var end_date string
	var start_dateAgo string
	var end_dateAgo string
	if err := ctx.ReadJSON(&strFilter); err != nil {
		model.FailWithDetailed(err, "Failed to parse JSON Body", ctx)
		return
	}
	switch strFilter.Filter {
	case Today:
		start_date, end_date = services.GetStartAndEndTimeForToday()
		start_dateAgo, end_dateAgo = services.GetYesterday(start_date, end_date)
	case ThisMonth:
		start_date, end_date = services.GetStartAndEndTimeForThisMonth()
		start_dateAgo, end_dateAgo = services.Get30DaysAgo(start_date, end_date)
	case ThisYear:
		start_date, end_date = services.GetStartAndEndTimeForThisYear()
		start_dateAgo, end_dateAgo = services.GetLastYearStartAndEnd(start_date, end_date)
	default:
		model.FailWithMessage("Out of range request", ctx)
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
		"sales":    sales,
		"increase": percentage,
		"status":   increaseOrDecrease,
	}

	model.OkWithData(data, ctx)
}

func (d *dashboardcontroller) GetRevenue(ctx iris.Context) {
	//var filter model.Filter
	var strFilter model.StrFilter
	var start_date string
	var end_date string
	var start_dateAgo string
	var end_dateAgo string
	if err := ctx.ReadJSON(&strFilter); err != nil {
		model.FailWithDetailed(err, "Failed to parse JSON Body", ctx)
		return
	}
	switch strFilter.Filter {
	case Today:
		start_date, end_date = services.GetStartAndEndTimeForToday()
		start_dateAgo, end_dateAgo = services.GetYesterday(start_date, end_date)
	case ThisMonth:
		start_date, end_date = services.GetStartAndEndTimeForThisMonth()
		start_dateAgo, end_dateAgo = services.Get30DaysAgo(start_date, end_date)
	case ThisYear:
		start_date, end_date = services.GetStartAndEndTimeForThisYear()
		start_dateAgo, end_dateAgo = services.GetLastYearStartAndEnd(start_date, end_date)
	default:
		model.FailWithMessage("Out of range request", ctx)
		return
	}

	revenues, err := services.DashboardService.GetRevenuesService(start_date, end_date)
	if err != nil {
		model.FailWithDetailed(err, "Failed get Revenues data", ctx)
		return
	}

	percentage, increaseOrDecrease := services.DashboardService.GetRevenueIncreaseDecrease(start_date, end_date, start_dateAgo, end_dateAgo)
	if math.IsNaN(percentage) {
		percentage = float64(0)
	}
	data := iris.Map{
		"revenues": revenues,
		"increase": percentage,
		"status":   increaseOrDecrease,
	}

	model.OkWithData(data, ctx)
}

func (d *dashboardcontroller) GetCustomers(ctx iris.Context) {
	//
	var filter model.Filter
	if err := ctx.ReadJSON(&filter); err != nil {
		model.FailWithDetailed(err, "Failed to parse JSON Body", ctx)
		return
	}

	customers, err := services.DashboardService.GetCustomersService(filter)
	if err != nil {
		model.FailWithDetailed(err, "Failed get Revenues data", ctx)
		return
	}
	model.OkWithData(customers, ctx)
}

func (d *dashboardcontroller) GetRecentSales(ctx iris.Context) {
	//
	var strFilter model.StrFilter
	if err := ctx.ReadJSON(&strFilter); err != nil {
		model.FailWithDetailed(err, "Failed to parse JSON Body", ctx)
		return
	}
	switch strFilter.Filter {
	case "Today":
		start_time, end_time := services.GetStartAndEndTimeForToday()
		fmt.Println(start_time)
		fmt.Println(end_time)
	case "This month":
		start_time, end_time := services.GetStartAndEndTimeForThisMonth()
		fmt.Println(start_time)
		fmt.Println(end_time)
	case "This year":
		start_time, end_time := services.GetStartAndEndTimeForThisYear()
		fmt.Println(start_time)
		fmt.Println(end_time)
	default:
		model.FailWithMessage("Out of range request", ctx)
	}
	//start_time, end_time := services.GetStartAndEndTimeForToday()
	//fmt.Println(start_time)
	//fmt.Println(end_time)
	//var count int64
	//util.Master().Model(model.Customers{}).Where("created_at BETWEEN ? AND ?", start_time, end_time).Count(&count)
	//fmt.Println(count, "==================")
}

func (d *dashboardcontroller) GetTopSelling(ctx iris.Context) {
	//
}
