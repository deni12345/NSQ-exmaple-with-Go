package main

import (
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	decodeConfig := nsq.NewConfig()
	c, err := nsq.NewConsumer("Sample_Topic", "Sample_Channel", decodeConfig)
	if err != nil {
		log.Panic("Could not create consumer")
	}

	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println("NSQ message received")
		log.Printf("value %v", message.Body)
		return nil
	}))

	err = c.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	log.Println("Awaiting messages from NSQ topic \"Sample_Topic\"...")
	wg.Wait()
}
