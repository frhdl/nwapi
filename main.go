package main

import (
	"fmt"
	"log"

	"github.com/getchipman/nwapi/models"
	"github.com/getchipman/nwapi/pkg/setting"
)

func main() {
	setting.Setup()
	models.Setup()

	result, err := models.GetEmployees()
	if err != nil {
		log.Fatalln("Erro:", err)
	}

	fmt.Println(result[0].LastName)
}
