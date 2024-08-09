package main

import (
	"context"
	"fmt"
	"os"
	"todo/internal/repository"
	"todo/pkg/datastore/sqlite"
)

func main() {
	db, err := sqlite.Open()
	if err != nil {
		fmt.Println(err)
	}
	repo := repository.New(db)
	todos, err := repo.List(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(todos)
}
