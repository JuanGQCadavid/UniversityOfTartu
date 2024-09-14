package generators

import (
	"hw2/internal/domain"
	"log"
	"strings"
)

func GenerateTokens(operation string) *domain.Operation {
	operations := strings.Split(operation, " ")

	if len(operations) < 3 {
		log.Println("Leng is less that expedted for operation ", operation)
		return nil
	}
	queueStackID := operations[2]

	if len(operations) == 3 {
		return &domain.Operation{
			DataType: domain.OperationDataType(queueStackID[0]),
			DatatId:  queueStackID[1:],
			OpType:   domain.REMOVE,
			Timing:   operations[0],
		}
	}

	if len(operations) == 4 {
		return &domain.Operation{
			DataType: domain.OperationDataType(queueStackID[0]),
			DatatId:  queueStackID[1 : len(queueStackID)-1], //Avoiding the ,
			OpType:   domain.INSERT,
			Id:       operations[3],
			Timing:   operations[0],
		}
	}

	return &domain.Operation{}
}
