package generators

import (
	"fmt"
	"hw2/internal/domain"
	"hw2/internal/domain/ports"
	"os"

	"github.com/gocarina/gocsv"
)

const (
	generalCSVFileName = "cmd/ex1/results/hw1_general.csv"
	allCSVFileName     = "cmd/ex1/results/hw1_all.csv"
)

// Generates two independents csv files, one with general information (Total op, total errs, total insert/delete)
// And a second csv with more details per data structure
func GenerateCSV(dataStructures map[domain.OperationDataType]map[string]ports.DataStructure, totalReports int) {

	allStats, general := GetAllStats(dataStructures)
	writeCSVToFile(allStats, allCSVFileName)
	writeCSVToFile([]domain.GeneralCSV{general}, generalCSVFileName)
}

func GenerateMessagesStatsReport(sqsNumber, stackNumber int, msgStats map[string]*domain.MsgStat) {
	fileName := fmt.Sprintf("%d_sqs_%d_stack_report.csv", sqsNumber, stackNumber)
	dataToPrint := make([][]string, 0, len(msgStats)+1)

	dataToPrint[0] = []string{
		"Data id",
		"Data type",
		"Mean delta time",
		"Oldest id",
		"Oldest created time",
		"Oldest removed time",
		"Youngest id",
		"Youngest created time",
		"Youngest removed time",
	}

	for _, msg := range msgStats {
		dataToPrint = append(dataToPrint, []string{
			msg.DataId,
			string(msg.DataType),
			fmt.Sprintf("%04f", msg.MeanDeltaTime),
			msg.Oldest.Id,
			msg.Oldest.Metadata.TimeCreated,
			msg.Oldest.Metadata.TimeDeleted,
			msg.Youngest.Id,
			msg.Youngest.Metadata.TimeCreated,
			msg.Youngest.Metadata.TimeDeleted,
		})

	}

	StringListToCSV(dataToPrint, fileName)

}

func writeCSVToFile(data any, fileName string) {
	gFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}

	if err = gocsv.MarshalFile(data, gFile); err != nil {
		panic(err.Error())
	}
}

func StringListToCSV(data [][]string, fileNme string) {
	// PENDING
}
