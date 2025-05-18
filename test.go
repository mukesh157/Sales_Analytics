package main

import (
	// "SalesAnalytcs/global"
	global "SalesAnalytcs/Global"
	"log"
	"time"
)

type EligibleForVote struct {
	IsEligible string
	Status     string
	ErrorMsg   string
}

func main() {

	log.Println("Main(+)")

	var lResp EligibleForVote
	lDOB := "12/04/2003"
	lResp.Status = global.SuccessCode
	lIsEligible, lErr := IsELigibleForVote(lDOB)
	if lErr != nil {
		log.Println("Error M001", lErr)
		lResp.Status = "E"
		lResp.ErrorMsg = lErr.Error()
	}
	log.Println("lIsEligible :", lIsEligible)

	log.Println("Main(-)")
}

func IsELigibleForVote(pDOB string) (string, error) {
	var lIsEligible string
	lDOB, lErr := time.Parse("02/01/2003", pDOB)
	if lErr != nil {
		log.Println("Error IEFV001", lErr)
		return lIsEligible, lErr
	}
	lCurrentDate := time.Now()
	lCurrentYear := lCurrentDate.Year()

	lAge := lCurrentYear - lDOB.Year()

	if lAge >= 18 {
		lIsEligible = "Y"
	} else {
		lIsEligible = "N"
	}
	return lIsEligible, nil
}
