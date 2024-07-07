package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)



func main(){
  go func() {
      log.Println("Consumer started")
      if err := http.ListenAndServe(":3030", nil); err != nil {
        log.Fatalf("HTTP server failed: %v", err)
      }
    }()
  
  config := sarama.NewConfig() 

  consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
  if err != nil {
    panic(err)
  }

  defer func() {
    if err := consumer.Close(); err != nil {
      log.Fatalln(err)
    }
  }()

  partitionConsumer, err := consumer.ConsumePartition("TestTopic", 0, sarama.OffsetNewest)
  if err != nil {
    panic(err)
  }

  defer func() {
    if err := partitionConsumer.Close(); err != nil {
      log.Fatalln(err)
    }
  }()

  // Trap SIGINT to trigger a shutdown.
  signals := make(chan os.Signal, 1)
  signal.Notify(signals, os.Interrupt)
// Consume messages in a loop
	go func() {
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				log.Printf("Consumed message: %s, offset: %d", string(msg.Value), msg.Offset)
        if string(msg.Value) == "shutdown"{
          os.Exit(0)
        }
			case err := <-partitionConsumer.Errors():
				log.Printf("Error consuming messages: %v", err)
			case sig := <-signals:
				if sig != nil {
					log.Println("Signal received, shutting down:", sig)
					return
				}
			}
		}
	}()
// Wait for a signal to exit
	sig := <-signals
	if sig != nil {
		log.Println("Signal received, shutting down:", sig)
	}

}

