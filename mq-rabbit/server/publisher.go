package server

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func SendMsg(conection *ConectionMQ, message interface{}) {

	conectionURI := fmt.Sprintf("amqp://%s:%s@%s:%s/", conection.Username, conection.Password, conection.Host, conection.Port)
	conn, err := amqp.Dial(conectionURI)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		conection.QueueName, // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body, err := json.Marshal(message)
	failOnError(err, "Failed to marshal JSON")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Timestamp:   time.Now(),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")

}
