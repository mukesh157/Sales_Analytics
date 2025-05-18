package revenue

import (
	"SalesAnalytcs/DbConn"
	global "SalesAnalytcs/Global"
	"SalesAnalytcs/Sales_Analytics/common"
	"SalesAnalytcs/Sales_Analytics/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"strings"
)

func Get_revenuedetails(w http.ResponseWriter, r *http.Request) {

	log.Println("Get_revenuedetails (+) ")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Indicator, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, credentials")

	if r.Method == http.MethodPost {

		lIndicator := r.Header.Get("Indicator")
		var linputRec model.GetDetails
		var lGetRevenueDetailsRec model.GetRevenueDetails
		lGetRevenueDetailsRec.Status = global.SuccessCode

		lErr := json.NewDecoder(r.Body).Decode(&linputRec)
		if lErr != nil {
			log.Println("Error in decoding GRD01 ", lErr)
			return
		}

		if strings.TrimSpace(linputRec.StartDate) == "" || strings.TrimSpace(linputRec.EndDate) == "" {
			fmt.Errorf("date field mandatory GRD02")
			return
		}

		switch lIndicator {
		case "Date_range":
			lGetRevenueDetailsRec.Total_revenue, lErr = GetTotalRevenue(linputRec)

		case "Product":
			lGetRevenueDetailsRec.TotProdRevenue, lErr = GetRevenuebyProd(linputRec)

		case "Category":
			lGetRevenueDetailsRec.TotalcatRevenue, lErr = GetRevenuebyCat(linputRec)

		case "Region":
			lGetRevenueDetailsRec.TotalRevenue_byreg, lErr = GetRevenuebyregion(linputRec)
		}

		if lErr != nil {

			return
		}

		lData, lErr := json.Marshal(lGetRevenueDetailsRec)
		if lErr != nil {
			log.Println("Error GRD03 ", lErr)
			return
		}

		fmt.Fprint(w, string(lData))

	}

	log.Println("Get_revenuedetails (-) ")

}

func GetTotalRevenue(pInputRec model.GetDetails) (string, error) {
	var ltotalRevenue string

	lStartDate := common.GetDate(pInputRec.StartDate)
	lEndDate := common.GetDate(pInputRec.EndDate)

	lErr := DbConn.GlobalDB.GormDb.Debug().Table("orders AS o").
		Select("nvl(SUM((oi.quantity_sold * oi.unit_price * (1 - oi.discount)) + o.shipping_cost),'') AS total_revenue").
		Joins("JOIN order_items AS oi ON o.id = oi.order_id").
		Where("o.date_of_sale BETWEEN ? AND ?", lStartDate, lEndDate).
		Scan(&ltotalRevenue).Error
	if lErr != nil {
		log.Println("Error GTR001", lErr)
		return ltotalRevenue, lErr
	}

	return ltotalRevenue, nil
}

func GetRevenuebyProd(pInputRec model.GetDetails) ([]model.ProductRevenue, error) {

	log.Println(" GetRevenuebyProd (+) ")

	var ltotalRevenue_byprod []model.ProductRevenue

	lStartDate := common.GetDate(pInputRec.StartDate)
	lEndDate := common.GetDate(pInputRec.EndDate)

	lErr := DbConn.GlobalDB.GormDb.Debug().Table("order_items").
		Select("products.name, nvl(SUM(order_items.quantity_sold  * order_items.unit_price * (1 - order_items.discount)),'') as total_revenue").
		Joins("JOIN products ON order_items.product_id = products.id").
		Joins("JOIN orders ON order_items.order_id  = orders.id").
		Where("orders.date_of_sale BETWEEN ? AND ?", lStartDate, lEndDate).
		Group("order_items.product_id, products.name").
		Scan(&ltotalRevenue_byprod).Error

	if lErr != nil {
		log.Println("Error GRBP01", lErr)
		return ltotalRevenue_byprod, lErr
	}

	log.Println(" GetRevenuebyProd (-) ")
	return ltotalRevenue_byprod, nil
}

func GetRevenuebyCat(pInputRec model.GetDetails) ([]model.CategoryRevenue, error) {

	log.Println("GetRevenuebyCat(+)")

	var ltotalRevenue_bycat []model.CategoryRevenue

	lStartDate := common.GetDate(pInputRec.StartDate)
	lEndDate := common.GetDate(pInputRec.EndDate)

	lErr := DbConn.GlobalDB.GormDb.Debug().Table("order_items").
		Select("products.category as category, nvl(SUM(order_items.quantity_sold  * order_items.unit_price * (1 - order_items.discount)),'') as total_revenue").
		Joins("JOIN products ON order_items.product_id = products.id").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("orders.date_of_sale BETWEEN ? AND ?", lStartDate, lEndDate).
		Group("products.category").
		Scan(&ltotalRevenue_bycat).Error

	if lErr != nil {
		log.Println("Error GRBC01", lErr)
		return ltotalRevenue_bycat, lErr
	}

	log.Println("GetRevenuebyCat(-)")
	return ltotalRevenue_bycat, nil
}

func GetRevenuebyregion(pInputRec model.GetDetails) ([]model.RegionRevenue, error) {

	log.Println(" GetRevenuebyregion (+) ")

	var ltotalRevenueByregion []model.RegionRevenue

	lStartDate := common.GetDate(pInputRec.StartDate)
	lEndDate := common.GetDate(pInputRec.EndDate)

	lErr := DbConn.GlobalDB.GormDb.Table("order_items").
		Select("nvl(orders.region,''), nvl(SUM(order_items.quantity_sold  * order_items.unit_price * (1 - order_items.discount)),'') as total_revenue").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("orders.date_of_sale BETWEEN ? AND ?", lStartDate, lEndDate).
		Group("orders.region").
		Scan(&ltotalRevenueByregion).Error

	if lErr != nil {
		log.Println("Error GRBR01", lErr)
		return ltotalRevenueByregion, lErr
	}

	log.Println(" GetRevenuebyregion (-) ")
	return ltotalRevenueByregion, nil
}
