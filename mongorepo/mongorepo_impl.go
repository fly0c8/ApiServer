package mongorepo

import (
	"github.com/fly0c8/ApiServer/model"
	"log"
)

type MongoContractRepository struct {}
func(r *MongoContractRepository) SaveContract(contract *model.ContractModel) bool {
	log.Println("Saving contract to mongodb...")
	return true
}