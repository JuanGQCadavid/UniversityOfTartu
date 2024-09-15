package main

import (
	"fmt"
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
	PYTHON_FILE_NAME    = "repositories/base_python_code.py"
	PYTHON_FILE_NAME_v2 = "repositories/base_python_code.py"
	BASE_FOLDER_OUTPUTS = "cmd/ex2/results/hw2_"
	generalCSVFileName  = "cmd/ex2/results/hw2_general.csv"
	allCSVFileName      = "cmd/ex2/results/hw2_all.csv"
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
	cmd := exec.Command("python3", PYTHON_FILE_NAME_v2, strconv.Itoa(QUEUE_NUMBER), strconv.Itoa(STACK_NUMBER))
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalln(err.Error())
	}

	response := strings.Split(string(out), "\n")
	response = response[:len(response)-1] // Deleting last enter

	historial := make([]*domain.Msg, 0)

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
			if msg, err := ds.Pop(op.Timing); err == nil {
				historial = append(historial, msg)
			}
		}
	}

	lastOperation := generators.GenerateTokens(response[len(response)-1])
	fmt.Println(lastOperation.Timing)
	before := len(historial)

	initTiming, err := strconv.Atoi(lastOperation.Timing[1:])

	if err != nil {
		panic(err)
	}

	for {
		before := len(historial)
		initTiming++

		timeFormarted := fmt.Sprintf("t%d", initTiming)
		fmt.Println(timeFormarted, initTiming)
		for datatype := range dataStructures {
			for _, datastructure := range dataStructures[datatype] {
				if val, err := datastructure.Pop(timeFormarted); err == nil {
					historial = append(historial, val)
				}
			}
		}

		if before == len(historial) { // If it does not increase in one iteration it means that there were not more messages leff
			break
		}

	}

	fmt.Println(before, len(historial))
	generators.GenerateMessagesStatsReport(QUEUE_NUMBER, STACK_NUMBER, BASE_FOLDER_OUTPUTS, generators.MessagesStats(historial))
	generators.GenerateCSV(dataStructures, len(response)-1, generalCSVFileName, allCSVFileName)
}
