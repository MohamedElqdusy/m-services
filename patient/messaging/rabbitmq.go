package messaging

import (
	"log"

	"patient/utils"

	"github.com/streadway/amqp"
)

type RabbitMqStore struct {
	conn *amqp.Connection
}

func NewRabbitMqStore(url string) (*RabbitMqStore, error) {
	if url == "" {
		panic("Empty RabbitMq connection Url")
	}
	conn, err := amqp.Dial(url)
	if err != nil {
		panic("Failed to connect to RabbitMq at: " + url)
	}
	return &RabbitMqStore{conn: conn}, nil

}

func (m *RabbitMqStore) Publish(body []byte, exchangeName string, routingName string) error {
	if m.conn == nil {
		panic("RabbitMq connection is not initialized.")
	}
	ch, err := m.conn.Channel()
	// TODO: channel closing
	//defer ch.Close()

	err = ch.Publish(
		exchangeName, // exchange
		routingName,  // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	log.Printf(" [x] Sent %s to to exchange %s with routing key %s", body, exchangeName, routingName)
	return err
}

func (m *RabbitMqStore) Subscribe(exchangeName string, routingName string, queueName string, consumerName string, handlerFunc func(amqp.Delivery)) error {
	ch, err := m.conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	// TODO: channel closing
	//defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	utils.FailOnError(err, "Failed to declare an exchange")

	queue, err := ch.QueueDeclare(
		queueName, // name of the queue
		false,     // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	utils.FailOnError(err, "Failed to declare an Queue")

	////////////////////////////////////////////
	//////////////////////////////////////
	log.Printf("Binding queue %s to exchange %s with routing key %s",
		queue.Name, exchangeName, routingName)
	err = ch.QueueBind(
		queue.Name,   // queue name
		routingName,  // routing key
		exchangeName, // exchange
		false,
		nil)
	utils.FailOnError(err, "Failed to bind a queue")
	////////////////////////////////////////////
	////////////////////////////////////////////

	msgs, err := ch.Consume(
		queue.Name,   // queue
		consumerName, // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	go receiveMesagesPeriodically(msgs, handlerFunc)
	return nil
}

func (m *RabbitMqStore) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}

func receiveMesagesPeriodically(msgs <-chan amqp.Delivery, handlerFunc func(d amqp.Delivery)) {
	for msg := range msgs {
		handlerFunc(msg)
	}
}
