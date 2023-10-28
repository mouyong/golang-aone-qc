package rabbitmq

import (
	"context"
	"log"
	"time"
    "fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitmqClient *amqp.Connection
var RabbitmqChannel *amqp.Channel

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func NewRabbitmq(host string, port int) {
    amqpHost := fmt.Sprintf("amqp://guest:guest@%s:%d/", host, port)
    conn, err := amqp.Dial(amqpHost)
    failOnError(err, "Failed to connect to RabbitMQ")
    RabbitmqClient = conn

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    RabbitmqChannel = ch
}

func Send(queue_name string, data string) {
    q, err := RabbitmqChannel.QueueDeclare(
        queue_name, // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")
    
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    body := data
    err = RabbitmqChannel.PublishWithContext(ctx,
    "",     // exchange
    q.Name, // routing key
    false,  // mandatory
    false,  // immediate
    amqp.Publishing {
        ContentType: "text/plain",
        Body:        []byte(body),
    })
    failOnError(err, "Failed to publish a message")
    log.Printf(" [x] Sent %s\n", body)
}

func Close() {
    RabbitmqChannel.Close()
    RabbitmqClient.Close()
}