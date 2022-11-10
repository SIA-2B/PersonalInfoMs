//Consumer

package commons

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

type personaInput struct {
	IdPersona string `json:"idPersona"`
}

func RabbitMQConsumer() {
	ip := "172.20.0.2"

	conn, err := amqp.Dial("amqp://grupo-2b:123456789@" + ip + ":5672")

	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("Conexión con RabbitMQ establecida en " + ip)
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
	/*
		go func() {
			for deliery := range chDelibery {

				a := deliery.Body
				fmt.Println("msg: " + string(a))
			}
		}()
	*/

	noStop := make(chan bool)

	go func() {
		for deliery := range chDelibery {

			a := string(deliery.Body)
			log.Println("msg: " + a)

			//str := `{"idPersona":"6"}`
			res := personaInput{}
			json.Unmarshal([]byte(a), &res)
			log.Println(res.IdPersona)

			if err != nil {
				log.Println(err)
				return
			}

			//Producer Thread
			go func() {

				//Consulta DB
				r := consultaDBExistPerson(res.IdPersona)

				log.Println("{\"idPersona\": " + res.IdPersona + ",\"volver\": " + strconv.FormatBool(r) + "}")

				err = ch.Publish(
					"",
					"direct",
					false,
					false,
					amqp.Publishing{
						Headers:     nil,
						ContentType: "text/plain",
						Body:        []byte("{\"idPersona\": \"" + res.IdPersona + "\",\"volver\": " + strconv.FormatBool(r) + "}"),
					})

				if err != nil {
					log.Println(err)
					return
				}

				time.Sleep(2 * time.Second)
			}()
		}
	}()

	<-noStop
}

func consultaDBExistPerson(id string) bool {

	personaID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Invalid ID")
		return false
	}

	db := ConexionDB()
	rows, err := db.Query("SELECT COUNT(*) FROM Persona WHERE idPersona = ? ;", personaID)

	defer rows.Close()

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Println(err)
		}
	}
	resultado := false

	if count >= 1 {
		resultado = true
	} else {
		resultado = false
	}

	return resultado
}
