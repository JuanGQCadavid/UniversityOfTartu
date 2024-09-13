package main

import (
	"fmt"
	"hw2/internal/domain"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const (
	PYTHON_FILE_NAME = "repositories/base_python_code.py"
)

var (
	QUEUE_NUMBER = 5
	STACK_NUMBER = 3
)

func main() {
	cmd := exec.Command("python3", PYTHON_FILE_NAME, strconv.Itoa(QUEUE_NUMBER), strconv.Itoa(STACK_NUMBER))
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalln(err.Error())
	}

	response := strings.Split(string(out), "\n")
	for i := 0; i < 10; i++ {
		fmt.Printf("Original: %s\nToken: %+v\n\n", response[i], extractOperation(response[i]))

	}
}

func extractOperation(operation string) *domain.Operation {
	operations := strings.Split(operation, " ")
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
