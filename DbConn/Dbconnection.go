package DbConn

import (
	"database/sql"
	"fmt"
	"log"

	// helperpkg "sales_analytics/helper_pkg"

	// helperpkg "sales_analytics/helper_pkg"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Dbconnection() (*gorm.DB, *sql.DB, error) {
	log.Println("Dbconnection (+) ")

	var lConstring string
	var lDb *sql.DB
	var lErr error
	var lGormDb *gorm.DB

	lDbDetailsRec := Db_detailReading()

	lConstring = `` + lDbDetailsRec.Mysql.User + `:` + lDbDetailsRec.Mysql.Password + `@tcp(` + lDbDetailsRec.Mysql.Server + `:` + fmt.Sprint(lDbDetailsRec.Mysql.Port) + `)/` + lDbDetailsRec.Mysql.Database + `?charset=utf8mb4&parseTime=True&loc=Local`
	lGormDb, lErr = gorm.Open(mysql.Open(lConstring), &gorm.Config{})
	if lErr != nil {
		log.Println(lErr)
		return lGormDb, lDb, lErr
	}

	lDb, lErr = lGormDb.DB()
	if lErr != nil {
		log.Println(lErr)
		return lGormDb, lDb, lErr
	}

	lConpoolconfig := connectionpoolConfig()

	lDb.SetMaxIdleConns(lConpoolconfig.OpenConnCount)

	lDb.SetMaxOpenConns(lConpoolconfig.IdleConnCount)

	lDb.SetConnMaxIdleTime(time.Second * time.Duration(lConpoolconfig.MaxIdleConnCount))

	log.Println("Dbconnection (-) ")
	return lGormDb, lDb, lErr
}
