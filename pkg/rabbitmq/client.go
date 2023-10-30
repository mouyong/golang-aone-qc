package rabbitmq

import (
	"aone-qc/internal/handlers"
	"aone-qc/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

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

func StartQcQueue() {
	start_qc_queue_name := "start_qc_queue"
	ch := RabbitmqChannel

	q, err := ch.QueueDeclare(
		start_qc_queue_name, // name
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			// log.Printf("Received a message: %s", d.Body)
			// dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(1)
			time.Sleep(t * time.Second)
			log.Printf("Done")

			var taskDTO map[string]interface{}
			var qcTaskDTO handlers.QcTaskDTO

			json.Unmarshal(d.Body, &taskDTO)

			fmt.Println(taskDTO["qc_task_samples"])

			qc_tasksBytes, _ := json.Marshal(taskDTO["qc_tasks"])
			json.Unmarshal(qc_tasksBytes, &qcTaskDTO)

			qcTaskModel := models.QcTasks{
				Environment:       qcTaskDTO.Environment,
				TenantID:          qcTaskDTO.TenantId,
				Slug:              qcTaskDTO.Slug,
				ProjectName:       qcTaskDTO.ProjectName,
				ExperimentBatchNo: qcTaskDTO.ExperimentBatchNo,
				AnalysesBatchNo:   qcTaskDTO.AnalysesBatchNo,
			}

			qcTaskId, err := qcTaskModel.Save()
			if err != nil {
				fmt.Println("save error: ", err)
				return
			}

			var qc_task_samples []*models.QcTaskSample
			for _, sample := range taskDTO["qc_task_samples"].([]interface{}) {
				fmt.Println(sample)

				sampleItem, _ := sample.(map[string]interface{})

				qc_task_sample := &models.QcTaskSample{
					BatchID:           qcTaskId,
					ExperimentBatchNo: sampleItem["experiment_batch_number"].(string),
					AnalysesBatchNo:   sampleItem["analysis_batch_number"].(string),
					SampleNo:          sampleItem["test_number"].(string),
				}
				qc_task_sample.Save()

				qc_task_samples = append(qc_task_samples, qc_task_sample)
			}

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func ListenQueue() {
	go StartQcQueue()
}

func Send(queue_name string, data string) {
	ch := RabbitmqChannel

	q, err := ch.QueueDeclare(
		queue_name, // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := data
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
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
