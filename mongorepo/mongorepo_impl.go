package mongorepo

import (
	"log"

	"github.com/fly0c8/ApiServer/model"
)

type MongoContractRepository struct{}

func (r *MongoContractRepository) SaveContract(contract *model.ContractModel) bool {
	log.Println("Saving contract to mongodb...")
	return true
}
