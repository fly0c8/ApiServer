package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/fly0c8/ApiServer/contractservice"
	"github.com/nats-io/nats.go"
)

type Server struct{
	nc *nats.Conn
}

func (s *Server) AddContract(ctx context.Context, req *contractservice.AddContractReq) (*contractservice.AddContractRes, error) {

	s.nc.Publish("config", []byte(req.ContractId))
	res := &contractservice.AddContractRes{Success:true}
	return res, nil
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


	_, err = nc.Subscribe("config", func(m *nats.Msg) {
		log.Printf("%s\n", m.Data)
	})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully subscribed to config")
	}

	srv := &Server{nc:nc}
	twirphandler := contractservice.NewContractServiceServer(srv, nil)
	http.ListenAndServe(":8080", twirphandler)

}
