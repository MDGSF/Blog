package models

import (
	"fmt"

	"github.com/MDGSF/Blog/setting"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"

	// import mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
documents at here: http://gorm.io/docs/
*/

var gDB *gorm.DB

// DBStart start db connection
func DBStart() {
	var err error
	sqlSetting := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		setting.DBUser, setting.DBPassWord, setting.DBHost, setting.DBPort, setting.DBName)
	gDB, err = gorm.Open("mysql", sqlSetting)
	if err != nil {
		beego.Error("failed to open mysql", err)
		return
	}
	//defer gDB.Close()

	beego.Info("open mysql success", setting.DBUser, setting.DBPassWord, setting.DBHost, setting.DBPort, setting.DBName)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "tb_" + defaultTableName
	}

	// auto generate table.
	gDB.AutoMigrate(&User{})
}
