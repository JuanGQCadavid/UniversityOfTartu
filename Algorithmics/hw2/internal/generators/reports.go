package generators

import (
	"encoding/csv"
	"fmt"
	"hw2/internal/domain"
	"hw2/internal/domain/ports"
	"os"

	"github.com/gocarina/gocsv"
)

// Generates two independents csv files, one with general information (Total op, total errs, total insert/delete)
// And a second csv with more details per data structure
func GenerateCSV(dataStructures map[domain.OperationDataType]map[string]ports.DataStructure, totalReports int, generalCSVFileName, allCSVFileName string) {

	allStats, general := GetAllStats(dataStructures)
	writeCSVToFile(allStats, allCSVFileName)
	writeCSVToFile([]domain.GeneralCSV{general}, generalCSVFileName)
}

func GenerateMessagesStatsReport(sqsNumber, stackNumber int, baseFolder string, msgStats map[string]*domain.MsgStat) {
	fileName := fmt.Sprintf("%d_sqs_%d_stack_report.csv", sqsNumber, stackNumber)
	dataToPrint := make([][]string, len(msgStats)+1)

	dataToPrint[0] = []string{
		"Data id",
		"Data type",
		"Mean delta time",
		"Oldest id",
		"Oldest created time",
		"Oldest removed time",
		"Oldest delta time",
		"Youngest id",
		"Youngest created time",
		"Youngest removed time",
		"Youngest delta time",
	}
	index := 1

	for _, msg := range msgStats {
		dataToPrint[index] = []string{
			msg.DataId,
			string(msg.DataType),
			fmt.Sprintf("%.2f", msg.MeanDeltaTime),
			msg.Oldest.Id,
			msg.Oldest.Metadata.TimeCreated,
			msg.Oldest.Metadata.TimeDeleted,
			fmt.Sprint(msg.Oldest.Metadata.DeltaTime),
			msg.Youngest.Id,
			msg.Youngest.Metadata.TimeCreated,
			msg.Youngest.Metadata.TimeDeleted,
			fmt.Sprint(msg.Youngest.Metadata.DeltaTime),
		}
		index++

	}

	StringListToCSV(dataToPrint, baseFolder, fileName)

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

func StringListToCSV(data [][]string, baseFolder, fileNme string) {
	f, err := os.OpenFile(fmt.Sprint(baseFolder, fileNme), os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}

	writter := csv.NewWriter(f)

	for _, record := range data {
		if err := writter.Write(record); err != nil {
			panic(err)
		}
	}

	writter.Flush()
}
