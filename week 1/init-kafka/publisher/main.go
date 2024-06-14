package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/IBM/sarama"
)



func main(){
  go func() {
      log.Println("Publisher started")
      if err := http.ListenAndServe(":3000", nil); err != nil {
        log.Fatalf("HTTP server failed: %v", err)
      }
    }()
  config := sarama.NewConfig() 
  producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
  if err != nil {
    panic(err)
  }

  defer func() {
    if err := producer.Close(); err != nil {
      log.Fatalln(err)
    }
  }()

  // Trap SIGINT to trigger a shutdown.
  inputScanner := make(chan string, 1)
  fmt.Println("Please type a message to publish. Type shutdown to end publishing and consuming")
  for {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := strings.TrimSpace(scanner.Text())
    inputScanner<-input
    select {
    case message:=<-inputScanner:
      producer.Input() <- &sarama.ProducerMessage{Topic: "TestTopic", Partition: 0, Value: sarama.StringEncoder(input)}
      if message == "shutdown"{
        os.Exit(0)
      }
      case err := <-producer.Errors():
        log.Println("Failed to produce message:", err)
    }
  }
}
