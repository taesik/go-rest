package main

import (
	"context"
	"fmt"
	"go-rest/internal/db"
)

func Run() error {
	fmt.Println("starting up ")

	database, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to create database")
		return err
	}
	if err := database.Ping(context.Background()); err != nil {
		return err
	}

	fmt.Println("Successfully connected and pinged database")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
