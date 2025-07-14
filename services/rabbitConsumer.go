package services

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

func StartRabbitConsumer() {
	conn, err := amqp.Dial("amqp://rabbitmq")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"appointments", // name
		"fanout",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %s", err)
	}

	q, err := ch.QueueDeclare(
		"",    // empty name creates a random queue
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	err = ch.QueueBind(
		q.Name,        // queue name
		"",            // routing key
		"appointments", // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind queue: %s", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		true,   // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %s", err)
	}

	go func() {
		for d := range msgs {
			var event map[string]interface{}
			json.Unmarshal(d.Body, &event)
			log.Printf("📥 Received AppointmentCreated event: %+v", event)
			// aquí podrías mandar a InfluxDB con influxService
		}
	}()

	log.Printf(" [*] Waiting for AppointmentCreated events. To exit press CTRL+C")
	select {}
}
