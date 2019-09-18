package main

import (
	"log"
	"github.com/streadway/amqp"
	"strconv"
	"time"
	"strings"
	"math"
	"fmt"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func StdDev(numbers []float64, mean float64) float64 {
    total := 0.0
    for _, number := range numbers {
        total += math.Pow(number-mean, 2)
    }
    variance := total / float64(len(numbers)-1)
    return math.Sqrt(variance)
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	sq, serr := ch.QueueDeclare(
		"to-server", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(serr, "Failed to declare Server queue")

	cq, cerr := ch.QueueDeclare(
		"to-client",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(serr, "Failed to declare Client queue")

	msgs, cerr := ch.Consume(
		cq.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(cerr, "Failed to register a consumer")

	times := make([]int, 10000)
	timesFloated := make([]float64, 10000)
	forever := make(chan bool)
	i := 0

	go func() {
		for d := range msgs {
			s1 :=  strings.Split(string(d.Body), "..")
			t, err := strconv.Atoi(s1[0])
			n, err := strconv.Atoi(s1[1])
			if err == nil {
			}		
			times[n] = int(time.Now().UTC().UnixNano()) - t
			timesFloated[n] = float64(float64(time.Now().UTC().UnixNano()) - float64(t))
			i++
			if (i >= 9999){
				forever <- false
			}
		}
	}()
		
	for i:=0;i<10000;i++{
		nsec := int(time.Now().UTC().UnixNano())
		body := (strconv.Itoa(nsec)) + ".." + (strconv.Itoa(i))
		serr = ch.Publish(
			"",     // exchange
			sq.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(serr, "Failed to publish a message")
	}

	<- forever
	sum := 0
	for i:=range times{
		sum += times[i]
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
