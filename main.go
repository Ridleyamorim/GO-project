package main

import (
	"database/sql"

	"github.com/ridleyamorim/go-intensivo/internal/infra/database"
	"github.com/ridleyamorim/go-intensivo/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Model string
	Color string
}

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)

	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput {
		ID: "123",
		Price: 10.0,
		Tax: 1.0,
	}

	output, err := uc.Execute(input)

	if err != nil {
		panic(err)
	}
	println(output)
}
