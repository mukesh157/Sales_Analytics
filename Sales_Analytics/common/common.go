package common

import (
	// helperpkg "sales_analytics/helper_pkg"
	"log"
	"time"
)

func GetDate(pdateStr string) time.Time {

	lDate, lErr := time.Parse("2006-01-02", pdateStr)
	if lErr != nil {
		log.Println("Error GD001", lErr)
	}

	return lDate
}
