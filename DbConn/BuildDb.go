package DbConn

import "log"

func BuildConnection() error {
	log.Println("BuildConnection (+) ")
	var lErr error

	GlobalDB.GormDb, GlobalDB.Db, lErr = Dbconnection()
	if lErr != nil {
		log.Println("Error (DCBC01) ", lErr.Error())
		return lErr
	}

	log.Println("BuildConnection (-) ")
	return lErr
}
