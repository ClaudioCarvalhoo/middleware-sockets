package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"
	"github.com/streadway/amqp"
	"fmt"
	"math"
)

func StdDev(numbers []float64, mean float64) float64 {
    total := 0.0
    for _, number := range numbers {
        total += math.Pow(number-mean, 2)
    }
    variance := total / float64(len(numbers)-1)
    return math.Sqrt(variance)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func fibonacciRPC(n int) (res int, err error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	corrId := randomString(32)
	err = ch.Publish(
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(strconv.Itoa(n)),
		})
	failOnError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {
			break
		}
	}

	return
}

func main() {
	n := 2
	times := make([]time.Duration, 10000)
	timesFloated := make([]float64, 10000)

	for i:=0;i<10000;i++{
		log.Printf(" [x] Requesting fib(%d)", n)
		start := time.Now()
		res, err := fibonacciRPC(n)
		failOnError(err, "Failed to handle RPC request")
		elapsed := time.Since(start)
		times[i] = elapsed
		timesFloated[i] = float64(elapsed.Nanoseconds())
		log.Printf(" [.] Got %d", res)
	}
	sum := int64(0)
	for _, t := range times {
		sum += t.Nanoseconds()
	}
	mean := sum / 10000
	std := StdDev(timesFloated, float64(mean))
	fmt.Print("Mean: ")
	fmt.Print(mean)
	fmt.Println(" nanoseconds")
	fmt.Print("Stdev: ")
	fmt.Print(std)
	fmt.Println(" nanoseconds")

}