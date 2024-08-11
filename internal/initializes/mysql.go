package initializes

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/phongnd2802/go-ecommerce/global"
	"go.uber.org/zap"
	_ "github.com/go-sql-driver/mysql"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func initMysql() {
	m := global.Config.MySQL

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DbName)
	db, err := sql.Open("mysql", s)
	checkErrorPanic(err, "InitMysql Error")
	global.Logger.Info("InitMysql Successfully")
	global.Db = db

	setPool()

}

func setPool() {
	m := global.Config.MySQL

	global.Db.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns) * time.Second)
	global.Db.SetMaxOpenConns(m.MaxOpenConns)
	global.Db.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
}