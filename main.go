package main

import (
	"fmt"

	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dns := os.GetEnv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connect Successfull.")
	}
	fmt.Println(db)
}
