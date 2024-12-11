package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/etf1/opentelemetry-go-contrib/instrumentation/github.com/confluentinc/confluent-kafka-go/otelconfluent"
	"demo"
	"net/http"
	"io"
)

func main() {

	confluentProducer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka-1:9092"})
	if err != nil {
		panic(err)
	}

	// Initialize OpenTelemetry trace provider and wrap the original kafka producer.
	tracerProvider := demo.InitTracer()
	p := otelconfluent.NewProducerWithTracing(confluentProducer, otelconfluent.WithTracerProvider(tracerProvider))
	defer p.Close()

	// HTTP Server
	h := Hello{producer: p}
	http.HandleFunc("/producer", h.sendDataToKafka)
	http.ListenAndServe(":8888", nil)

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
}

type Hello struct {
	producer *otelconfluent.Producer
}

func ( h *Hello )sendDataToKafka(w http.ResponseWriter, r *http.Request) {
	// Produce messages to topic (asynchronously)
	topic := "myTopic"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		h.producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}
	h.producer.Flush(15 * 1000)
	io.WriteString(w, "Hello, HTTP!\n")
}
