package main

import (
	"context"
	"github.com/fly0c8/ApiServer/model"
	"github.com/fly0c8/ApiServer/mongorepo"
	"github.com/fly0c8/ApiServer/postgresrepo"
	"log"
	"net/http"
	"time"

	"github.com/fly0c8/ApiServer/contractservice"
	"github.com/nats-io/nats.go"
)


type ContractRepository interface {
	SaveContract(contract* model.ContractModel) bool
}

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

	mongoRepo := &mongorepo.MongoContractRepository{}
	postgresRepo := &postgresrepo.PostgresContractRepository{}

	_, err = nc.Subscribe("config", func(m *nats.Msg) {
		log.Printf("%s\n", m.Data)
		contract := &model.ContractModel{
			UUID:       "",
			ContractId: "",
		}
		mongoRepo.SaveContract(contract)
		postgresRepo.SaveContract(contract)
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
