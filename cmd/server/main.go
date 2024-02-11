package main

import (
	"fmt"

	"github.com/anil1226/go-employee/internal/db"
	"github.com/anil1226/go-employee/internal/employee"
	transhttp "github.com/anil1226/go-employee/internal/transport/http"
)

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}

func Run() error {
	fmt.Println("starting up")
	db, err := db.NewDatabase()
	if err != nil {
		return err
	}

	serv := employee.NewService(db)
	httpHandler := transhttp.NewHandler(serv)

	if err = httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}
