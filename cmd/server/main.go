package main

import "fmt"

func Run() error {
	fmt.Println("starting up ")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
