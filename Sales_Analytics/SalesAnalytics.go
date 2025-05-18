package salesanalyticsprocess

import (
	global "SalesAnalytcs/Global"
	"SalesAnalytcs/Sales_Analytics/ReadFile"
	"SalesAnalytcs/Sales_Analytics/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// global "sales_analytics/Global"
	// helperpkg "sales_analytics/helper_pkg"
	// "sales_analytics/sales_analyticsprocess/model"
	// "sales_analytics/sales_analyticsprocess/readfile"
	// "sales_analytics/tomlreader"
)

func UploadFileDetails(w http.ResponseWriter, r *http.Request) {
	log.Println("UploadFileDetails(+)")
	var lRespRec model.Response
	lRespRec.Status = global.SuccessCode
	lRespRec.Status = "File uploaded"

	lErr := ReadCSVFile()
	if lErr != nil {
		log.Println("Error (UFD001)", lErr)
		fmt.Fprint(w, lErr)
		return
	}

	lData, lErr := json.Marshal(lRespRec)
	if lErr != nil {
		log.Println("Error (UFD002)", lErr)
		fmt.Fprint(w, lErr)
		return
	}

	fmt.Fprint(w, string(lData))
	log.Println("UploadFileDetails(-)")

}

func ReadCSVFile() error {
	log.Println("ReadCSVFile (+) ")

	lErr := ReadFile.CSV_Reader("./CSVFile/saleanalytics.csv")
	if lErr != nil {
		log.Println("Error (RFUD001)", lErr)
		return lErr
	}

	log.Println("ReadCSVFile (-) ")
	return nil

}
