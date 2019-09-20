package models

import (
// mysql dialects
// _ "github.com/jinzhu/gorm/dialects/mysql"
)

// var db *gorm.DB

// SetUp init database
func SetUp() {
	// var err error
	// db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	setting.DatabaseSetting.User,
	// 	setting.DatabaseSetting.Password,
	// 	setting.DatabaseSetting.Host,
	// 	setting.DatabaseSetting.Name))
	// if err != nil {
	// 	log.Fatalf("models.Setup err: %v", err)
	// }

	// // db.SingularTable(true)
	// db.DB().SetMaxIdleConns(10)
	// db.DB().SetMaxOpenConns(100)

}

// CloseDB close database
// func CloseDB() {
// 	defer db.Close()
// }
