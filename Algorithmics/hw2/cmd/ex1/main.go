package main

import (
	"hw2/internal/domain"
	"hw2/internal/domain/ports"
	"hw2/internal/generators"
	"hw2/internal/queues"
	"hw2/internal/stacks"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const (
	PYTHON_FILE_NAME   = "repositories/base_python_code.py"
	generalCSVFileName = "cmd/ex1/results/hw1_general.csv"
	allCSVFileName     = "cmd/ex1/results/hw1_all.csv"
)

var (
	QUEUE_NUMBER                                                               = 5
	STACK_NUMBER                                                               = 3
	sqsDS          map[string]ports.DataStructure                              = make(map[string]ports.DataStructure)
	stackDS        map[string]ports.DataStructure                              = make(map[string]ports.DataStructure)
	dataStructures map[domain.OperationDataType]map[string]ports.DataStructure = make(map[domain.OperationDataType]map[string]ports.DataStructure)
)

func init() {
	// Crating hash maps
	for i := 1; i <= STACK_NUMBER; i++ {
		stackDS[strconv.Itoa(i)] = stacks.NewStack(strconv.Itoa(i))
	}

	for i := 1; i <= QUEUE_NUMBER; i++ {
		sqsDS[strconv.Itoa(i)] = queues.NewQueue(strconv.Itoa(i))
	}
	dataStructures[domain.QUEUE] = sqsDS
	dataStructures[domain.STACK] = stackDS
}

func main() {
	cmd := exec.Command("python3", PYTHON_FILE_NAME, strconv.Itoa(QUEUE_NUMBER), strconv.Itoa(STACK_NUMBER))
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalln(err.Error())
	}

	response := strings.Split(string(out), "\n")

	for _, strOp := range response {
		op := generators.GenerateTokens(strOp)
		if op == nil {
			continue
		}
		ds := dataStructures[op.DataType][op.DatatId]

		if ds == nil {
			log.Fatal("Data type ", op.DataType, " with id ", op.DatatId, " Does not exits on index for string ", strOp)
		}
		if op.OpType == domain.INSERT {
			ds.Push(op.Id, op.Timing)
		}

		if op.OpType == domain.REMOVE {
			ds.Pop(op.Timing)
		}

	}
	generators.GenerateCSV(dataStructures, len(response)-1, generalCSVFileName, allCSVFileName)
}
