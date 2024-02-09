package main

import (
	"fmt"

	"github.com/anil1226/go-employee/internal/db"
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

	emp, err := db.ReadItem()
	if err != nil {
		return err
	}
	fmt.Printf("%+v", emp)
	// if err = db.Ping(context.Background()); err != nil {
	// 	return err
	// }

	return nil
}
