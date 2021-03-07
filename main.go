package main

import (
	"golang-crowdfunding-backend/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:ReviRzA89@!@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Test Save from Service"
	userInput.Email = "example@gmail.com"
	userInput.Occupation = "musician"
	userInput.Password = "password"

	userService.RegisterUser(userInput)

}
