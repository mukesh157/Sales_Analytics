package main

import (
	dbconnection "SalesAnalytcs/DbConn"
	salesanalyticsprocess "SalesAnalytcs/Sales_Analytics"
	"SalesAnalytcs/Sales_Analytics/products"
	"SalesAnalytcs/Sales_Analytics/revenue"
	"SalesAnalytcs/tomlreader"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server Started")

	lFile, lErr := os.OpenFile("./log/logfile"+time.Now().Format("02012006.15.04.05.000000000")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if lErr != nil {
		log.Fatalf("Error Opening File: %v", lErr)
	}
	defer lFile.Close()

	log.SetOutput(lFile)

	lErr = dbconnection.BuildConnection()
	if lErr != nil {
		log.Fatalf("Error connecting Database :%v", lErr)
	}
	defer dbconnection.GlobalDB.Db.Close()

	// lAutorunConfig := tomlreader.ReadTomlFile("./toml/serviceconfig.toml")
	go DaiyFileUploadProcess()

	lRouter := mux.NewRouter()
	lRouter.HandleFunc("/Uploadfiledetails", salesanalyticsprocess.UploadFileDetails).Methods(http.MethodGet)
	lRouter.HandleFunc("/based_onRevenue", revenue.Get_revenuedetails).Methods(http.MethodPost)
	lRouter.HandleFunc("/Get_Productsdetails", products.Get_Productsdetails).Methods(http.MethodPost)

	lSrv := &http.Server{
		Handler: lRouter,
		Addr:    ":23434",
	}

	log.Fatal(lSrv.ListenAndServe())
}

func DaiyFileUploadProcess() {
	for {

		lNow := time.Now()

		lTimeConfig := tomlreader.ReadTomlFile("./toml/Config.toml")
		lHour := fmt.Sprintf("%v", lTimeConfig.(map[string]interface{})["Hour"])
		lminute := fmt.Sprintf("%v", lTimeConfig.(map[string]interface{})["Minute"])

		lHour_int, lErr := strconv.Atoi(lHour)
		if lErr != nil {
			lHour_int = 8
		}
		lminute_int, lErr := strconv.Atoi(lminute)
		if lErr != nil {
			lHour_int = 0
		}

		if lNow.Hour() == lHour_int && lNow.Minute() == lminute_int {

			lErr := salesanalyticsprocess.ReadCSVFile()
			if lErr != nil {
				log.Fatalf("Error ", lErr)
			}

			time.Sleep(61 * time.Second)

		} else {
			time.Sleep(30 * time.Second)
		}

	}
}
