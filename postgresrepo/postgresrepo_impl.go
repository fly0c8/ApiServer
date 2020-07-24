package postgresrepo

import (
	"github.com/fly0c8/ApiServer/model"
	"log"
)

type PostgresContractRepository struct {}
func(r *PostgresContractRepository) SaveContract(contract *model.ContractModel) bool {
	log.Println("Saving contract to postgres...")
	return true
}