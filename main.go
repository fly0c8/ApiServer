package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/fly0c8/ApiServer/contractservice"

	"github.com/nats-io/nats.go"
)

type Server struct{}

func (s *Server) AddContract(context.Context, *contractservice.AddContractReq) (*contractservice.AddContractRes, error) {
	return nil, nil

}

func main() {
	nc, err := nats.Connect(
		nats.DefaultURL,
		nats.PingInterval(20*time.Second),
		nats.MaxPingsOutstanding(5))

	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	log.Println("Subscribing...")
	wg := sync.WaitGroup{}
	wg.Add(1)

	_, err = nc.Subscribe("config", func(m *nats.Msg) {
		log.Printf("%s\n", m.Data)
	})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully subscribed to config")
	}

	go func() {
		for {
			now := time.Now().Unix()
			nc.Publish("config", []byte(fmt.Sprintf("Unix time in seconds: %v", now)))
			time.Sleep(time.Second * 5)
		}

	}()

	wg.Wait()

}
