package main

import (
	"fmt"

	"config/app/controllers"
	"config/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
