package main

import (
	"github.com/IBM/sarama"
	"log"
	"strconv"
	"time"
)

func Produce(topic string, limit int) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewSyncProducer([]string{"10.7.6.12:30001"}, config)
	if err != nil {
		log.Fatal("NewSyncProducer err:", err)
	}
	defer producer.Close()
	for i := 0; i < limit; i++ {
		str := strconv.Itoa(int(time.Now().UnixNano()))

		msg := &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(str)}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Println("SendMessage err: ", err)
			return
		}
		log.Printf("[Producer] partitionid: %d; offset:%d, value: %s\n", partition, offset, str)
	}
}
func Consume(topic string) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"10.7.6.12:30001"}, config)
	if err != nil {
		log.Fatal("NewConsumer err: ", err)
	}
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("ConsumePartition err: ", err)
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		log.Printf("[Consumer] partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, string(message.Value))
	}
}
func main() {
	Produce("test", 100000)
	//Consume("test")
	//Produce("test", 1000)
}
