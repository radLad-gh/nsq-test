package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

type Message struct {
	Name      string
	Content   string
	Timestamp string
}

func Produce() {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}
	topic := "Topic_Example"
	msg := Message{
		Name:      "Message Name Example",
		Content:   "Message Content Example",
		Timestamp: time.Now().String(),
	}
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	err = producer.Publish(topic, payload)
	if err != nil {
		log.Println(err)
	}
}
