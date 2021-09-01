package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	deliveryChan := make(chan kafka.Event)

	producer := NewKafkaProducer()
	//Publish("mensagem02", "teste", producer, nil, deliveryChan)
	Publish("transferiu", "teste", producer, []byte("transferencia"), deliveryChan)

	go DeliveryReport(deliveryChan) //async

	fmt.Println("Jonathan")
	producer.Flush(1000)

	/*
		e := <-deliveryChan
		msg := e.(*kafka.Message)
		if msg.TopicPartition.Error != nil {
			fmt.Println("Erro ao enviar")
		} else {
			fmt.Println("Mensagem enviada:", msg.TopicPartition)
		}
		producer.Flush(1000)
	*/
	//fmt.Println("Hello Go")
}

func NewKafkaProducer() *kafka.Producer {

	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "gokafka_kafka_1:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",  //"0, 1 ou all"
		"enable.idempotence":  "true", //"false, true"
		//"bootstrap.servers": "gokafka_kafka_1:9094",
	}

	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}

	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar")
			} else {
				fmt.Println("Mensagem enviada:", ev.TopicPartition)
				// anotar no banco de dados que a mensagem foi processado.
				// ex: confima que uma transferencia bancaria ocorreu.
			}
		}
	}
}
