package main

import (
	"fmt"
	"golang_test/config"
	"golang_test/infrastructure"
	"golang_test/infrastructure/registry"
	"log"

	"github.com/labstack/echo"
)

func main() {
	config.ReadConfig()

	db := infrastructure.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = infrastructure.NewRouter(e, r.NewAppController())

	port := config.C.Server.Address
	fmt.Println("Server listen at http://localhost" + ":" + port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalln(err)
	}
}
