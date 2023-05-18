package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

type messageHandler struct{}

func Consume() {
	config := nsq.NewConfig()
	config.MaxAttempts = 10
	config.MaxInFlight = 5
	config.MaxRequeueDelay = time.Second * 900
	config.DefaultRequeueDelay = time.Second * 0
	topic := "Topic_Example"
	channel := "Channel_Example"
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal(err)
	}
	consumer.AddHandler(&messageHandler{})
	consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	consumer.Stop()
}

func (h *messageHandler) HandleMessage(m *nsq.Message) error {
	var request Message
	if err := json.Unmarshal(m.Body, &request); err != nil {
		log.Println("Error when Unmarshaling the message body, Err : ", err)

		return err
	}
	log.Println("Message")
	log.Println("--------------------")
	log.Println("Name : ", request.Name)
	log.Println("Content : ", request.Content)
	log.Println("Timestamp : ", request.Timestamp)
	log.Println("--------------------")
	log.Println("")
	return nil
}
