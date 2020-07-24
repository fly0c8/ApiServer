package contractservice

import (
	context "context"

	"github.com/nats-io/nats.go"
)

type Server struct {
	nc *nats.Conn
}

func NewServer(nc *nats.Conn) *Server {
	return &Server{nc: nc}
}

func (s *Server) AddContract(ctx context.Context, req *AddContractReq) (*AddContractRes, error) {
	s.nc.Publish("config", []byte(req.ContractId))
	res := &AddContractRes{Success: true}
	return res, nil
}
