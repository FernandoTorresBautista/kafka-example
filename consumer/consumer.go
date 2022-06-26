package consumer

import (
	"fmt"

	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
)

// TOPIC to consume messages
var TOPIC = "comments"
var KAFKA_CONN = "localhost:9092"

// InitConsumer: proccess the messages on the topic = TOPIC
func InitConsumer() {
	// Get the connection to Fafa
	worker, err := ConnectConsumer([]string{KAFKA_CONN})
	if err != nil {
		panic("ERROR CONNETION KAFKA:" + err.Error())
	}
	// Calling ConsumePartition. It will open one connection per broker
	// and share it for all partitions that live on it.
	consumer, err := worker.ConsumePartition(TOPIC, 0, sarama.OffsetOldest)
	// consumer, err := worker.ConsumePartition(TOPIC, 0, sarama.OffsetNewest)
	if err != nil {
		panic("ERROR OPEN CONNETION PER BROKER:" + err.Error())
	}
	fmt.Println("Consumer started...")
	// Make a signal to interrup
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// Count how many message processed
	msgCount := 0
	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n", msgCount, string(msg.Topic), string(msg.Value))
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
	fmt.Println("Processed", msgCount, "messages")
	if err := worker.Close(); err != nil {
		panic("ERROR CLOSE CONNECTION KAFKA :" + err.Error())
	}
}

// Connect to Apache Kafka using sarama as Consumer
func ConnectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	// NewConsumer creates a new consumer using the given broker addresses and configuration
	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
