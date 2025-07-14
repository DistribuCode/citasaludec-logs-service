package services

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

func StartRabbitConsumer() {
	// Conectar a RabbitMQ
	conn, err := amqp.Dial("amqp://rabbitmq")
	if err != nil {
		log.Fatalf("❌ Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	// Abrir canal
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("❌ Failed to open a channel: %s", err)
	}
	defer ch.Close()

	// Declarar el exchange 'appointments' tipo fanout
	err = ch.ExchangeDeclare(
		"appointments", // exchange name
		"fanout",       // type
		false,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatalf("❌ Failed to declare exchange: %s", err)
	}

	// Declarar el queue fijo y durable
	q, err := ch.QueueDeclare(
		"appointments-logs", // nombre fijo
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		log.Fatalf("❌ Failed to declare queue: %s", err)
	}

	// Bindear el queue al exchange fanout
	err = ch.QueueBind(
		q.Name,              // nombre del queue
		"",                  // routing key vacío (fanout)
		"appointments",      // exchange name
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("❌ Failed to bind queue: %s", err)
	}

	// Consumir mensajes del queue
	msgs, err := ch.Consume(
		q.Name, // nombre del queue
		"",     // consumer tag
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("❌ Failed to register consumer: %s", err)
	}

	// Procesar mensajes recibidos
	go func() {
		for d := range msgs {
			var event map[string]interface{}
			if err := json.Unmarshal(d.Body, &event); err != nil {
				log.Printf("⚠️ Failed to unmarshal event: %s", err)
				continue
			}
			log.Printf("📥 Received AppointmentCreated event: %+v", event)
			// aquí podrías enviar a InfluxDB
		}
	}()

	log.Printf("🚀 Waiting for AppointmentCreated events. To exit press CTRL+C")
	select {}
}
