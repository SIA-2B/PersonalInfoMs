//Consumer

package commons

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func RabbitMQConsumer() {
	conn, err := amqp.Dial("amqp://admin:local23@172.17.0.4:5672")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	chDelibery, err := ch.Consume(
		"employees",
		"",
		true,
		false,
		false,
		false,
		nil)

	go func() {
		for deliery := range chDelibery {

			a := deliery.Body
			fmt.Println("msg: " + string(a))
		}
	}()

	/*
		noStop := make(chan bool)

		go func() {
			for deliery := range chDelibery {

				a := deliery.Body
				fmt.Println("msg: " + string(a))
			}
		}()

		<-noStop*/
}
