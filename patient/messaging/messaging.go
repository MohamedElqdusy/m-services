package messaging

import "github.com/streadway/amqp"

type MessageStore interface {
	Publish(msg []byte, exchangeName string, routingName string) error
	Subscribe(exchangeName string, routingName string, queueName string, consumerName string, handlerFunc func(amqp.Delivery)) error
	Close()
}

//Message client Implmentation used to enforce satisfying the MessageStore interface
var m MessageStore

func SetMessageStore(messageClient MessageStore) {
	m = messageClient
}

func Publish(msg []byte, exchangeName string, routingName string) error {
	return m.Publish(msg, exchangeName, routingName)
}
func Subscribe(exchangeName string, routingName string, queueName string, consumerName string, handlerFunc func(amqp.Delivery)) error {
	return m.Subscribe(exchangeName, routingName, queueName, consumerName, handlerFunc)
}

func Close() {
	m.Close()
}
