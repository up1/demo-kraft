package main

import (
	"fmt"
	"time"
	"demo"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/etf1/opentelemetry-go-contrib/instrumentation/github.com/confluentinc/confluent-kafka-go/otelconfluent"
)

func main() {

	confluentConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka-1:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}
	// Initialize OpenTelemetry trace provider and wrap the original kafka consumer.
	tracerProvider := demo.InitTracer()
	c := otelconfluent.NewConsumerWithTracing(confluentConsumer, otelconfluent.WithTracerProvider(tracerProvider))

	err = c.SubscribeTopics([]string{"myTopic"}, nil)

	if err != nil {
		panic(err)
	}

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
