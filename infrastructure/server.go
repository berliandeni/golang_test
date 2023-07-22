package infrastructure

import (
	"fmt"
	"golang_test/config"
	"golang_test/infrastructure/registry"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo"
)

func RunServer() {
	config.ReadConfig()

	db := NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = NewRouter(e, r.NewAppController())

	port := config.C.Server.Address
	fmt.Println("Server listen at http://localhost" + ":" + port)

	go func() {
		e.Logger.Fatal(e.Start(":" + port))
		// if err := e.Start(":" + port); err != nil {
		// 	log.Fatalln(err)
		// }
	}()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	s := <-c
	log.Fatalf("process killed with signal : %v\n", s.String())
}
