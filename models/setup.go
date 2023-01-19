package models

import (
	"fmt"
	//"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	//db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/testing_db"))
	//if err != nil {
	//	fmt.Println("Gagal koneksi database")
	//}

	dsn := "host=103.13.207.142 user=postgres password=User@mis1 dbname=dbusers sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Println("Gagal koneksi database")
	}

	db.AutoMigrate(&User{})

	DB = db
}
