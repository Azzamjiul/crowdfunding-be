package main

import (
	"fmt"
	"log"

	"github.com/azzamjiul/bwastartup/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(host.docker.internal:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database is good")

	var users []user.User
	length := len(users)
	fmt.Println(length)

	db.Find(&users)

	length = len(users)
	fmt.Println(length)

	for _, user := range users {
		fmt.Println(user.Name)
		fmt.Println(user.Email)
	}
}
