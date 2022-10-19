//Consumer

package commons

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func RabbitMQConsumer() {

	conn, err := amqp.Dial("amqp://ndcontrerasr:1234@172.17.0.5:5672")

	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Println(err)
		return
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

	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		for deliery := range chDelibery {

			a := deliery.Body
			fmt.Println("msg: " + string(a))
		}
	}()

	noStop := make(chan bool)

	go func() {
		for deliery := range chDelibery {

			a := deliery.Body
			fmt.Println("msg: " + string(a))
		}
	}()

	<-noStop
}
