package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	"go-rest/internal/db"
)

func Run() error {
	fmt.Println("starting up our application")

	database, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := database.MigrateDB(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("Failed to migrate database")
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
