package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/testcontainers/testcontainers-go/modules/redpanda"
)

func main() {
	//request := testcontainers.ContainerRequest{
	//	AlwaysPullImage: true,
	//}

	// Create a Redpanda container
	redpandaContainer, err := redpanda.Run(context.Background(),
		"docker.redpanda.com/redpandadata/redpanda:v23.3.3",
		//redpanda.WithEnableSASL(),
		//redpanda.WithEnableKafkaAuthorization(),
		redpanda.WithAutoCreateTopics(),

		//redpanda.WithEnableWasmTransform(),
		//redpanda.WithBootstrapConfig("data_transforms_per_core_memory_reservation", 33554432),
		//redpanda.WithBootstrapConfig("data_transforms_per_function_memory_limit", 16777216),
		//redpanda.WithNewServiceAccount("superuser-1", "test"),
		//redpanda.WithNewServiceAccount("superuser-2", "test"),
		//redpanda.WithNewServiceAccount("no-superuser", "test"),
		//redpanda.WithSuperusers("superuser-1", "superuser-2"),
		//redpanda.WithEnableSchemaRegistryHTTPBasicAuth(),
	)
	if err != nil {
		log.Fatalf("Failed to start Redpanda container: %v", err)
	}

	//defer func() {
	//	if err := testcontainers.TerminateContainer(redpandaContainer); err != nil {
	//		log.Printf("failed to terminate container: %s", err)
	//	}
	//}()

	// Get the broker address
	broker, err := redpandaContainer.PortEndpoint(context.Background(), "9092", "")
	if err != nil {
		log.Fatalf("Failed to get broker address: %v", err)
	}
	fmt.Printf("Redpanda broker running at: %s\n", broker)

	// Set up Kafka writer and reader
	topic := "example-topic"
	createTopic(broker, topic)

	writeMessages(broker, topic)
	readMessages(broker, topic)
}

func createTopic(broker, topic string) {
	fmt.Printf("Creating topic: %s\n", topic)

	// to create topics when auto.create.topics.enable='true'
	_, err := kafka.DialLeader(context.Background(), "tcp", broker, topic, 0)
	if err != nil {
		panic(err.Error())
	}

	//conn, err := kafka.Dial("tcp", broker)
	//if err != nil {
	//	log.Fatalf("Failed to connect to broker: %v", err)
	//}
	//defer conn.Close()
	//
	////controller, err := conn.Controller()
	////if err != nil {
	////	panic(err.Error())
	////}
	////
	////controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	////if err != nil {
	////	panic(err.Error())
	////}
	////defer controllerConn.Close()
	//
	//err = conn.CreateTopics(kafka.TopicConfig{
	//	Topic:             topic,
	//	NumPartitions:     1,
	//	ReplicationFactor: 1,
	//})
	//if err != nil {
	//	log.Fatalf("Failed to create topic: %v", err)
	//}
	fmt.Printf("Topic %s created successfully\n", topic)
}

func writeMessages(broker, topic string) {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(broker),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond, // Adjust batch timeout if needed
	}
	defer writer.Close()

	messages := []kafka.Message{
		{Key: []byte("key1"), Value: []byte("value1")},
		{Key: []byte("key2"), Value: []byte("value2")},
		{Key: []byte("key3"), Value: []byte("value3")},
	}

	fmt.Printf("Writing messages to topic: %s\n", topic)
	err := writer.WriteMessages(context.Background(), messages...)
	if err != nil {
		log.Fatalf("Failed to write messages: %v", err)
	}
	fmt.Println("Messages written successfully")
}

func readMessages(broker, topic string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	defer reader.Close()

	fmt.Printf("Reading messages from topic: %s\n", topic)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	Read(reader, ctx)
	Read(reader, ctx)
	Read(reader, ctx)
}

func Read(reader *kafka.Reader, ctx context.Context) {
	fmt.Println("Offset before read: ", reader.Offset())
	fmt.Println("Lag before read: ", reader.Lag())

	msg, err := reader.ReadMessage(ctx)
	if err != nil {
		if ctx.Err() != nil {
			fmt.Println("Finished reading messages (timeout)")
			fmt.Println(ctx.Err())
		}
		log.Fatalf("Failed to read messages: %v", err)
	}
	fmt.Printf("Received: key=%s value=%s\n", string(msg.Key), string(msg.Value))

	fmt.Println("Offset after read: ", reader.Offset())
	fmt.Println("Lag after read: ", reader.Lag())
}
