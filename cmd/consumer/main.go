package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ricassiocosta/gointensivo/internal/order/infra/database"
	usecase "github.com/ricassiocosta/gointensivo/internal/order/useCase"
	"github.com/ricassiocosta/gointensivo/pkg/rabbitmq"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPriceUseCase{
		OrderRepository: repository,
	}

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	out := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, out)

	for msg := range out {
		var inputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}

		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		fmt.Println(outputDTO)
	}
}
