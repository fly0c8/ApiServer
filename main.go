package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fly0c8/ApiServer/model"
	"github.com/fly0c8/ApiServer/mongorepo"
	"github.com/fly0c8/ApiServer/postgresrepo"

	"github.com/fly0c8/ApiServer/contractservice"
	"github.com/nats-io/nats.go"
)

type ContractRepository interface {
	SaveContract(contract *model.ContractModel) bool
}

var (
	repos = []ContractRepository{
		&mongorepo.MongoContractRepository{},
		&postgresrepo.PostgresContractRepository{},
	}
)

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
		contract := &model.ContractModel{
			UUID:       "",
			ContractId: "",
		}
		for _, repo := range repos {
			repo.SaveContract(contract)
		}

	})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully subscribed to config")
	}

	srv := contractservice.NewServer(nc)
	twirphandler := contractservice.NewContractServiceServer(srv, nil)
	http.ListenAndServe(":8080", twirphandler)

}
