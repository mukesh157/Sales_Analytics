package ReadFile

import (
	"SalesAnalytcs/DbConn"
	"SalesAnalytcs/Sales_Analytics/model"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func CSV_Reader(pFilepath string) error {

	//Opening the file
	lFile, lErr := os.Open(pFilepath)
	if lErr != nil {
		log.Println("Error CFR001", lErr)
		return lErr
	}

	//Closing the file
	defer lFile.Close()

	// Create a new CSV reader
	lReader := csv.NewReader(lFile)

	//Reading all the rows in the file
	lRecords, lErr := lReader.ReadAll()
	if lErr != nil {
		log.Println("Error CFR002", lErr)
		return lErr
	}

	for lIdx, lRows := range lRecords {

		if lIdx == 0 {
			continue
		}

		var lCustomerRec model.Customerdetails
		var lProductsRec model.Products
		var lOrdersRec model.Orders
		var lOrder_itemsRec model.Order_items

		lCustomerRec.Customer_id = lRows[2]
		lCustomerRec.Customer_name = lRows[12]
		lCustomerRec.Customer_email = lRows[13]
		lCustomerRec.Customer_address = lRows[14]

		lCustomerId, lErr := CheckAlreadyPresent("customers", `customer_id = '`+lCustomerRec.Customer_id+`'`)
		if lErr != nil {
			log.Println("Error CFR003", lErr)
			return lErr
		}

		if lCustomerId == 0 {
			lErr = DbConn.GlobalDB.GormDb.Table("customers").Create(&lCustomerRec).Error
			if lErr != nil {
				log.Println("Error CFR004", lErr)
				return lErr
			}

			lCustomerId = lCustomerRec.ID

		}

		lProductsRec.Product_id = lRows[1]
		lProductsRec.Name = lRows[3]
		lProductsRec.Category = lRows[4]

		lproductId, lErr := CheckAlreadyPresent("products", `product_id = '`+lProductsRec.Product_id+`'`)
		if lErr != nil {
			log.Println("Error CFR005", lErr)
			return lErr
		}

		if lproductId == 0 {
			lErr = DbConn.GlobalDB.GormDb.Table("products").Create(&lProductsRec).Error
			if lErr != nil {
				log.Println("Error CFR006", lErr)
				return lErr
			}

			lproductId = lProductsRec.ID

		}

		lOrdersRec.Order_id = lRows[0]
		lOrdersRec.Customer_id = lCustomerId
		lOrdersRec.Region = lRows[5]
		lOrdersRec.Date_of_sale = lRows[6]
		lOrdersRec.Payment_method = lRows[11]
		lOrdersRec.Shipping_cost = lRows[10]

		lOrderId, lErr := CheckAlreadyPresent("orders", `order_id = '`+lOrdersRec.Order_id+`'`)
		if lErr != nil {
			log.Println("Error CFR007", lErr)
			return lErr
		}

		if lOrderId == 0 {
			lErr = DbConn.GlobalDB.GormDb.Table("orders").Create(&lOrdersRec).Error
			if lErr != nil {
				log.Println("Error CFR008", lErr)
				return lErr
			}
		} else {
			continue
		}

		lOrder_itemsRec.Order_id = lOrdersRec.ID
		lOrder_itemsRec.Product_id = lproductId
		lOrder_itemsRec.Quantity_sold, lErr = strconv.Atoi(lRows[7])
		if lErr != nil {
			log.Println("Error CFR010", lErr)
			return lErr
		}
		lOrder_itemsRec.Unit_price, lErr = strconv.ParseFloat(lRows[8], 64)
		if lErr != nil {
			log.Println("Error CFR011", lErr)
			return lErr
		}
		lOrder_itemsRec.Discount, lErr = strconv.ParseFloat(lRows[9], 64)
		if lErr != nil {
			log.Println("Error CFR012", lErr)
			return lErr
		}

		lErr = DbConn.GlobalDB.GormDb.Table("order_items").Create(&lOrder_itemsRec).Error
		if lErr != nil {
			log.Println("Error CFR0013", lErr)
			return lErr
		}
	}
	return nil
}

func CheckAlreadyPresent(pTableName, pWherecon string) (uint, error) {
	log.Println(" CheckAlreadyPresent (+) ")

	var lId uint

	lErr := DbConn.GlobalDB.GormDb.Table(pTableName).Select("id").Where(pWherecon).Scan(&lId).Error
	if lErr != nil {
		log.Println("Error CAP001", lErr)
		return lId, lErr
	}

	log.Println(" CheckAlreadyPresent (-) ")
	return lId, lErr
}
