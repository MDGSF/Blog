package db

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

// Start start db connection
func Start() {
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
	gDB.AutoMigrate(&TUser{})
}

// func testDB() {

// 	go func() {
// 		for i := 0; i < 10000; i++ {
// 			user := &TUser{}
// 			user.ID = uint64(i)
// 			user.Query()
// 			time.Sleep(time.Millisecond * 10)
// 		}
// 	}()

// 	go func() {
// 		for i := 0; i < 10000; i++ {
// 			user := &TUser{}
// 			user.ID = uint64(i)
// 			user.Query()
// 			time.Sleep(time.Millisecond * 10)
// 		}
// 	}()

// 	go func() {
// 		for i := 0; i < 10000; i++ {
// 			user := &TUser{}
// 			user.ID = uint64(i)
// 			user.Delete()
// 			time.Sleep(time.Millisecond * 10)
// 		}
// 	}()

// 	for i := 0; i < 10000; i++ {
// 		user := &TUser{}
// 		user.ID = uint64(i)
// 		user.UserName = "root" + strconv.Itoa(i)
// 		user.PassWord = "123456"
// 		user.Email = "email" + strconv.Itoa(i)
// 		user.PhoneNumber = "phone" + strconv.Itoa(i)
// 		user.Create()
// 		time.Sleep(time.Millisecond * 10)
// 	}

// 	beego.Info("test db finished")

// 	// user2 := &TUser{ID: 1}
// 	// user2.Query()
// 	// beego.Info(user2)

// 	// user2.UserName = "huangjian"
// 	// user2.Update()

// 	// user3 := &TUser{}
// 	// user3.ID = 3
// 	// user3.UserName = "huangjian"
// 	// user3.Create()

// 	// user4 := &TUser{}
// 	// user4.ID = 4
// 	// user4.UserName = "huangjian"
// 	// user4.Create()

// 	// user5 := &TUser{}
// 	// user5.ID = 5
// 	// user5.UserName = "huangjian"
// 	// user5.Create()

// 	// user6 := &TUser{}
// 	// user6.ID = 5
// 	// user6.UserName = "huangjian"
// 	// user6.Delete()
// }
