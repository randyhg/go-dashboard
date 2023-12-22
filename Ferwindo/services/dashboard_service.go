package services

import (
	"fmt"
	"go-dashboard/model"
	"go-dashboard/util"
	"strconv"
	"time"
)

type dashboardservice struct{}

var DashboardService = new(dashboardservice)
var Today = "today"
var ThisMonth = "this_month"
var ThisYear = "this_year"

func (d *dashboardservice) GetSalesService(startDate, endDate string) (sales int64, err error) {
	if err := util.Master().Model(model.Sales{}).Where("date BETWEEN ? AND ?", startDate, endDate).Count(&sales).Error; err != nil {
		fmt.Println(err)
		return 0, err
	}
	return sales, nil
}

func (d *dashboardservice) GetSalesIncreaseDecrease(startDate1, endDate1, startDate2, endDate2 string) (percentageChange float64, increaseOrDecrease string) {
	increaseOrDecrease = "decrease"
	var today int64
	util.Master().Table("sales").Where("date BETWEEN ? AND ?", startDate1, endDate1).Count(&today)

	var yesterday int64
	util.Master().Table("sales").Where("date BETWEEN ? AND ?", startDate2, endDate2).Count(&yesterday)
	if yesterday == 0 && today > 0 {
		percentageChange = 100
	} else {
		percentageChange = float64(today-yesterday)/float64(yesterday)*100 - 100
	}
	if percentageChange >= 0 {
		increaseOrDecrease = "increase"
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
	increaseOrDecrease = "decrease"
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
		increaseOrDecrease = "increase"
	}
	return percentageChange, increaseOrDecrease
}

func (d *dashboardservice) GetCustomersService(startDate, endDate string) (customers int64, err error) {
	if err := util.Master().Model(model.Customers{}).Where("created_at BETWEEN ? AND ?", startDate, endDate).Count(&customers).Error; err != nil {
		fmt.Println(err)
		return 0, err
	}
	return customers, nil
}

func (d *dashboardservice) GetCustomerIncreaseDecrease(startDate1, endDate1, startDate2, endDate2 string) (percentageChange float64, increaseOrDecrease string) {
	increaseOrDecrease = "decrease"
	var today int64
	util.Master().Model(model.Customers{}).Where("created_at BETWEEN ? AND ?", startDate1, endDate1).Count(&today)

	var yesterday int64
	util.Master().Model(model.Customers{}).Where("created_at BETWEEN ? AND ?", startDate2, endDate2).Count(&yesterday)
	if yesterday == 0 && today > 0 {
		percentageChange = 100
	} else {
		percentageChange = float64(today-yesterday)/float64(yesterday)*100 - 100
	}
	if percentageChange >= 0 {
		increaseOrDecrease = "increase"
	}
	return percentageChange, increaseOrDecrease
}

func (d *dashboardservice) GetTopSellingService(start_date, end_date string) ([]model.TopSelling, error) {
	var topSelling []model.TopSelling
	if err := util.Master().Table("sale_lines AS sl").
		Select("p.product_image AS preview, p.name AS name, sl.unit_price price, SUM(sl.quantity) AS sold, SUM(sl.subtotal) AS revenue").
		Joins("INNER JOIN products AS p ON p.id = sl.product_id").
		Joins("INNER JOIN sales AS s ON s.id = sl.sale_id").
		Where("s.date BETWEEN ? AND ?", start_date, end_date).
		Group("preview, name, price").
		Order("revenue DESC").
		Limit(5).
		Scan(&topSelling).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return topSelling, nil
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

func SwitchCase(strFilter string) (filterName, start_date, end_date, start_dateAgo, end_dateAgo string, err error) {
	switch strFilter {
	case Today:
		filterName = "Today"
		start_date, end_date = GetStartAndEndTimeForToday()
		start_dateAgo, end_dateAgo = GetYesterday(start_date, end_date)
	case ThisMonth:
		filterName = "This Month"
		start_date, end_date = GetStartAndEndTimeForThisMonth()
		start_dateAgo, end_dateAgo = Get30DaysAgo(start_date, end_date)
	case ThisYear:
		filterName = "This Year"
		start_date, end_date = GetStartAndEndTimeForThisYear()
		start_dateAgo, end_dateAgo = GetLastYearStartAndEnd(start_date, end_date)
	default:
		return "", "", "", "", "", fmt.Errorf("Out of range request")
	}
	return
}

func FormatCurrency(amount string) (string, error) {
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return "", err
	}

	amountFormatted := strconv.FormatFloat(amountFloat, 'f', -1, 64)

	result := "Rp " + addThousandSeparator(amountFormatted)

	return result, nil
}

func addThousandSeparator(s string) string {
	parts := splitByDecimalPoint(s)

	parts[0] = addCommas(parts[0])

	return combineWithDecimalPoint(parts)
}

func splitByDecimalPoint(s string) []string {
	parts := []string{"", ""}

	decimalPointIndex := len(s)
	for i, c := range s {
		if c == '.' || c == ',' {
			decimalPointIndex = i
			break
		}
	}

	parts[0] = s[:decimalPointIndex]
	if decimalPointIndex < len(s) {
		parts[1] = s[decimalPointIndex+1:]
	}

	return parts
}

func combineWithDecimalPoint(parts []string) string {
	if parts[1] != "" {
		return parts[0] + "." + parts[1]
	}
	return parts[0]
}

func addCommas(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return addCommas(s[:n-3]) + "," + s[n-3:]
}
