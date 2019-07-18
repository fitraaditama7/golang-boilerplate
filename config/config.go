package config

import (
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:123459@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	//db.AutoMigrate(structs.Person{})
	return db
}
