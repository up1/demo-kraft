# Workshop :: Apache Kafka

## Install [Kafka with GraalVM](https://hub.docker.com/r/apache/kafka-native) + PLAINTEXT mode
* Smaller image size (faster download time)
* Faster startup time
* Lower memory usage
```
$export IMAGE=apache/kafka:3.8.0
$docker compose up -d kafka-1
$docker compose up -d kafka-2
$docker compose up -d kafka-3
$docker compose ps
$docker compose logs --follow
```

Kafka UI
```
$docker compose up -d kafka-ui
```

Access to UI Kafka
* http://localhost:8080/


## Run sample app
* [Kafka client library](https://github.com/confluentinc/confluent-kafka-go)

Producer
```
$docker compose up -d producer --build
```
Send message to Kafka
* http://localhost:8888/producer

Consumer
```
$docker compose up -d consumer --build
```

## Monitoring KafKa Broker
* [Kafka Exporter](https://github.com/danielqsj/kafka_exporter)
* JMX Exporter
* Prometheus
* Grafana

Start Kafka Exporter
```
$docker compose up -d kafka-exporter
```

Access to Kafka Exporter metric
* http://localhost:9308/metrics
  * Topic
  * Consumer group

Start JMX Exporter
```
$docker compose up -d jmx-exporter-broker
```

Access to Kafka Exporter metric
* http://localhost:5556/metrics

Start Prometheus
```
$docker compose up -d prometheus
```

Access to Prometheus dashboard
* http://localhost:9090/

Start Grafana
```
$docker compose up -d grafana
```

Access to Grafana dashboard
* http://localhost:3000
  * username=admin
  * password=password

## Tracing your message
* Producer
* Consumer
* Opentelemetry
* Jarger

Start Jaeger
```
$docker compose up -d jaeger
```

Access to Jaeger dashboard
* http://localhost:16686


Start Otel Collector
```
$docker compose up -d otel-collector
```
