package services

import (
	"fmt"
	"go-dashboard/model"
	"go-dashboard/util"
	"time"
)

type dashboardservice struct{}

var DashboardService = new(dashboardservice)

func (d *dashboardservice) GetSalesService(startDate, endDate string) (sales int64, err error) {
	if err := util.Master().Model(model.Sales{}).Where("date BETWEEN ? AND ?", startDate, endDate).Count(&sales).Error; err != nil {
		fmt.Println(err)
		return 0, err
	}
	return sales, nil
}

func (d *dashboardservice) GetSalesIncreaseDecrease(startDate1, endDate1, startDate2, endDate2 string) (percentageChange float64, increaseOrDecrease string) {
	increaseOrDecrease = "Decrease"
	var today int64
	util.Master().Table("sales").Where("date BETWEEN ? AND ?", startDate1, endDate1).Count(&today)

	var yesterday int64
	util.Master().Table("sales").Where("date BETWEEN ? AND ?", startDate2, endDate2).Count(&yesterday)
	if yesterday == 0 && today > 0 {
		percentageChange = 100
	} else {
		percentageChange = float64(today-yesterday) / float64(yesterday) * 100
	}
	if percentageChange >= 0 {
		increaseOrDecrease = "Increase"
	}
	return percentageChange, increaseOrDecrease
}

func (d *dashboardservice) GetRevenuesService(startDate, endDate string) (revenues string, err error) {
	if err := util.Master().Model(model.Sales{}).Select("SUM(grand_total) AS revenues").Where("date BETWEEN ? AND ?", startDate, endDate).Scan(&revenues).Error; err != nil {
		fmt.Println(err)
		return "0", nil
	}
	return revenues, nil
}

func (d *dashboardservice) GetRevenueIncreaseDecrease(startDate1, endDate1, startDate2, endDate2 string) (percentageChange float64, increaseOrDecrease string) {
	increaseOrDecrease = "Decrease"
	var today int64
	util.Master().Model(model.Sales{}).Select("SUM(grand_total) AS revenues").Where("date BETWEEN ? AND ?", startDate1, endDate1).Scan(&today)

	var yesterday int64
	util.Master().Model(model.Sales{}).Select("SUM(grand_total) AS revenues").Where("date BETWEEN ? AND ?", startDate2, endDate2).Scan(&yesterday)
	if yesterday == 0 && today > 0 {
		percentageChange = 100
	} else {
		percentageChange = float64(today-yesterday) / float64(yesterday) * 100
	}
	if percentageChange >= 0 {
		increaseOrDecrease = "Increase"
	}
	return percentageChange, increaseOrDecrease
}

func (d *dashboardservice) GetCustomersService(filter model.Filter) (customers int64, err error) {
	if err := util.Master().Model(model.Customers{}).Where("created_at BETWEEN ? AND ?", filter.StartDate, filter.EndDate).Count(&customers).Error; err != nil {
		fmt.Println(err)
		return 0, err
	}
	return customers, nil
}

func (d *dashboardservice) GetRecentSalesService(filter model.Filter) {

}

func GetStartAndEndTimeForToday() (string, string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	startTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location()).Format("2006-01-02 15:04:05")
	return startTime, now
}

func GetStartAndEndTimeForThisMonth() (string, string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	startTime := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location()).Format("2006-01-02 15:04:05")
	return startTime, now
}

func GetStartAndEndTimeForThisYear() (string, string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	startTime := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.Now().Location()).Format("2006-01-02 15:04:05")
	return startTime, now
}

func GetYesterday(start_date, end_date string) (string, string) {
	startDate, _ := time.Parse("2006-01-02 15:04:05", start_date)
	endDate, _ := time.Parse("2006-01-02 15:04:05", end_date)
	startYesterday := startDate.AddDate(0, 0, -1)
	endYesterday := endDate.AddDate(0, 0, -1)
	return startYesterday.Format("2006-01-02 15:04:05"), endYesterday.Format("2006-01-02 15:04:05")
}

func Get30DaysAgo(start_date, end_date string) (string, string) {
	startDate, _ := time.Parse("2006-01-02 15:04:05", start_date)
	endDate, _ := time.Parse("2006-01-02 15:04:05", end_date)
	startThirtyDaysAgo := startDate.AddDate(0, 0, -30)
	endThirtyDaysAgo := endDate.AddDate(0, 0, -30)
	return startThirtyDaysAgo.Format("2006-01-02 15:04:05"), endThirtyDaysAgo.Format("2006-01-02 15:04:05")
}

func GetLastYearStartAndEnd(start_date, end_date string) (string, string) {
	startDate, _ := time.Parse("2006-01-02 15:04:05", start_date)
	endDate, _ := time.Parse("2006-01-02 15:04:05", end_date)
	startLastYear := startDate.AddDate(-1, 0, 0)
	endLastYear := endDate.AddDate(-1, 0, 0)
	return startLastYear.Format("2006-01-02 15:04:05"), endLastYear.Format("2006-01-02 15:04:05")
}
