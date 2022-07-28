package main

import (
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

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{
		Name:       "Agus",
		Occupation: "Occupation",
		Password:   "adminadmin",
		Email:      "Agus@email.com",
	}

	userService.RegisterUser(userInput)
}
