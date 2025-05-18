package products

import (
	"SalesAnalytcs/DbConn"
	global "SalesAnalytcs/Global"
	"SalesAnalytcs/Sales_Analytics/common"
	"SalesAnalytcs/Sales_Analytics/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// global "sales_analytics/Global"
	// "sales_analytics/DbConn"
	// helperpkg "sales_analytics/helper_pkg"
	// "sales_analytics/sales_analyticsprocess/common"
	// "sales_analytics/sales_analyticsprocess/model"
	"strconv"
	"strings"
)

func Get_Productsdetails(w http.ResponseWriter, r *http.Request) {

	log.Println(" Get_Productsdetails (+) ")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Indicator, NINDICATOR, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, credentials")

	if r.Method == http.MethodPost {

		lIndicator := r.Header.Get("Indicator")
		lNINDICATOR := r.Header.Get("NINDICATOR")

		var lLmit int
		var linputRec model.GetDetails
		var lGetRevenueDetailsRec model.GetRevenueDetails
		lGetRevenueDetailsRec.Status = global.SuccessCode

		lErr := json.NewDecoder(r.Body).Decode(&linputRec)
		if lErr != nil {
			log.Println("Error GPD001", lErr)
			return
		}

		if strings.TrimSpace(linputRec.StartDate) == "" || strings.TrimSpace(linputRec.EndDate) == "" {
			log.Println("Error GPD002", lErr)
			return
		}

		if lNINDICATOR == "" {
			lLmit = 10
		} else {
			lLmit, lErr = strconv.Atoi(lNINDICATOR)
			if lErr != nil {
				log.Println("Error GPD003", lErr)
				lLmit = 10
			}
		}

		switch lIndicator {
		case "Overall":
			lGetRevenueDetailsRec.TopProduct, lErr = GetOverallProducts(linputRec, lLmit)

		case "Category":
			lGetRevenueDetailsRec.TopCategory, lErr = GetTopCategories(linputRec, lLmit)

		case "Region":
			lGetRevenueDetailsRec.TopRegion, lErr = GetTopRegions(linputRec, lLmit)
		}

		if lErr != nil {
			log.Println("Error GPD003", lErr)
			return
		}

		lData, lErr := json.Marshal(lGetRevenueDetailsRec)
		if lErr != nil {
			log.Println("Error GPD004", lErr)
			return
		}

		fmt.Fprint(w, string(lData))

	}
	log.Println(" Get_Productsdetails (-) ")

}

func GetOverallProducts(pInputRec model.GetDetails, lLimit int) ([]model.TopProduct, error) {

	var lTopProdectArr []model.TopProduct

	lStartDate := common.GetDate(pInputRec.StartDate)
	lEndDate := common.GetDate(pInputRec.EndDate)

	lErr := DbConn.GlobalDB.GormDb.Table("order_items").
		Select("nvl(products.name,''), nvl(SUM(order_items.quantity_sold ),'') as total_quantity").
		Joins("JOIN products ON order_items.product_id = products.id").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("orders.date_of_sale BETWEEN ? AND ?", lStartDate, lEndDate).
		Group("products.name").
		Order("total_quantity DESC").
		Limit(lLimit).
		Scan(&lTopProdectArr).Error

	if lErr != nil {
		log.Println("Error GOP001", lErr)
		return lTopProdectArr, lErr
	}

	return lTopProdectArr, nil
}

func GetTopCategories(pInputRec model.GetDetails, pLimit int) ([]model.TopCategory, error) {

	log.Println(" GetTopCategories (+) ")

	var lTopCategoryArr []model.TopCategory
	lStartDate := common.GetDate(pInputRec.StartDate)
	lEndDate := common.GetDate(pInputRec.EndDate)

	lErr := DbConn.GlobalDB.GormDb.Table("order_items").
		Select("nvl(products.category,'') as category, nvl(SUM(order_items.quantity_sold),'') as total_quantity").
		Joins("JOIN products ON order_items.product_id = products.id").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("orders.date_of_sale BETWEEN ? AND ?", lStartDate, lEndDate).
		Group("products.category").
		Order("total_quantity DESC").
		Limit(pLimit).
		Scan(&lTopCategoryArr).Error

	if lErr != nil {
		log.Println("Error GTC001", lErr)
		return lTopCategoryArr, lErr
	}

	log.Println(" GetTopCategories (-) ")

	return lTopCategoryArr, nil
}

func GetTopRegions(pInputRec model.GetDetails, pLimit int) ([]model.TopRegion, error) {
	log.Println(" GetTopRegions (+) ")

	var lTopRegionArr []model.TopRegion
	lStartDate := common.GetDate(pInputRec.StartDate)
	lEndDate := common.GetDate(pInputRec.EndDate)

	lErr := DbConn.GlobalDB.GormDb.Table("order_items").
		Select("nvl(orders.region,'') as region, nvl(SUM(order_items.quantity_sold),'') as total_quantity").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("orders.date_of_sale BETWEEN ? AND ?", lStartDate, lEndDate).
		Group("orders.region").
		Order("total_quantity DESC").
		Limit(pLimit).
		Scan(&lTopRegionArr).Error

	if lErr != nil {
		log.Println("Error GTR001", lErr)
		return lTopRegionArr, lErr
	}

	log.Println(" GetTopRegions (-) ")
	return lTopRegionArr, lErr

}
