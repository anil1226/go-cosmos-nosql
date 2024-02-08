package main

import "fmt"

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}

func Run() error {
	fmt.Println("starting up")
	return nil
}
