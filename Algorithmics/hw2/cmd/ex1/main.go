package main

import (
	"fmt"
	"hw2/internal/domain"
	"hw2/internal/domain/ports"
	"hw2/internal/queues"
	"hw2/internal/stacks"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const (
	PYTHON_FILE_NAME = "repositories/base_python_code.py"
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
		op := extractOperation(strOp)
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

	totalOp := 0
	for ds := range dataStructures {
		for _, dsInstance := range dataStructures[ds] {
			stats := dsInstance.GetStats()
			totalOp += stats.DeleteCount + stats.InsertCount + stats.ErrorsCount
			fmt.Println(stats.ToString())
		}
	}

	fmt.Println(totalOp, len(response)-1)
}

func extractOperation(operation string) *domain.Operation {
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
