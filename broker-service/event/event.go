package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// declareExchange declares the logs_topic topic exchange on the given channel.
func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        //auto-detechted
		false,        // internal
		false,        // no-wait
		nil,          // argumets
	)
}

// declareRandomQueue declares a transient, exclusive queue with a server-generated name.
func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
}
