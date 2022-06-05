package main

import (
	"fmt"
	"log"

	// "config/app/controllers"
	"config/app/models"
)

func main() {

	// controllers.StartMainServer()
	user, _ := models.GetUserByEmail("test@example.com")
	fmt.Println(user)
	session, err := user.CreateSession()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(session)
	valid,_ :=session.CheckSession()
	fmt.Println(valid)
}
