package config

/*
Folder ini digunakan untuk menyimpan file atau paket yang berhubungan dengan konfigurasi aplikasi. Misalnya, file yang berisi pengaturan koneksi database, konfigurasi API, dan setelan lainnya yang dapat digunakan di berbagai bagian aplikasi.
*/

import (
	"fmt"
	"golang-mysql/pkg/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {

	config := utils.Init()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
