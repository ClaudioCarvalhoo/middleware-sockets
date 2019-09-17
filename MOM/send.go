package main

import (
	"log"
	"github.com/streadway/amqp"
	"strconv"
	"time"
	"strings"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func printTimes(times []int){
	for i:= range times{
		log.Printf(strconv.Itoa(times[i]))
	}
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
	forever := make(chan bool)
	i := 0

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			s1 :=  strings.Split(string(d.Body), "..")
			t, err := strconv.Atoi(s1[0])
			n, err := strconv.Atoi(s1[1])
			if err == nil {
			}		
			times[n] = int(time.Now().UTC().UnixNano()) - t
			i++
			if (i >= 9999){
				printTimes(times)
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
		log.Printf(" [x] Sent %s", body)
	}

	<- forever
}
