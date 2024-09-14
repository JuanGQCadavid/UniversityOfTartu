package generators

import (
	"fmt"
	"hw2/internal/domain"
	"hw2/internal/domain/ports"
	"sort"
	"strconv"

	"gonum.org/v1/gonum/stat"
)

func castTimes(i *domain.Msg) (int, int) {
	start, errStart := strconv.Atoi(i.Metadata.TimeCreated[1:])
	end, errEnd := strconv.Atoi(i.Metadata.TimeDeleted[1:])

	if errEnd != nil || errStart != nil {
		fmt.Println("Start: ", i.Metadata.TimeCreated)
		fmt.Println("End: ", i.Metadata.TimeDeleted)
		panic("Start or end are not able to be convert")

	}

	return start, end
}

func GetAllStats(dataStructures map[domain.OperationDataType]map[string]ports.DataStructure) ([]domain.GeneralStats, domain.GeneralCSV) {
	allStats := make([]domain.GeneralStats, 0)
	general := domain.GeneralCSV{}

	for ds := range dataStructures {
		for _, dsInstance := range dataStructures[ds] {
			stats := dsInstance.GetStats()
			allStats = append(allStats, stats)

			general.TotalOperations += stats.DeleteCount + stats.InsertCount + stats.ErrorsCount
			general.TotalErrs += stats.ErrorsCount
			general.TotalInsert += stats.InsertCount
			general.TotalRemove += stats.DeleteCount

		}
	}

	return allStats, general
}

func MessagesStats(messages []*domain.Msg) map[string]*domain.MsgStat {
	sort.Slice(messages, func(i, j int) bool {

		if messages[i].Metadata.DeltaTime == 0 {
			iStart, iEnd := castTimes(messages[i])
			messages[i].Metadata.DeltaTime = iEnd - iStart
		}

		if messages[j].Metadata.DeltaTime == 0 {
			jStart, jEnd := castTimes(messages[j])
			messages[j].Metadata.DeltaTime = jEnd - jStart
		}

		return messages[i].Metadata.DeltaTime > messages[j].Metadata.DeltaTime

	})

	var statsPerDataStructure map[string]*domain.MsgStat = make(map[string]*domain.MsgStat)
	var deltaValues map[string][]float64 = make(map[string][]float64, len(messages))

	for _, msg := range messages {

		dsId := fmt.Sprintf("%s%s", msg.Metadata.DSType, msg.Metadata.DSId)
		if statsPerDataStructure[dsId] == nil {
			statsPerDataStructure[dsId] = &domain.MsgStat{
				DataType: msg.Metadata.DSType,
				DataId:   msg.Metadata.DSId,
				Oldest:   msg,
				Youngest: msg,
			}
			arr := make([]float64, 0)
			arr = append(arr, float64(msg.Metadata.DeltaTime))
			deltaValues[dsId] = arr
			continue
		}

		deltaValues[dsId] = append(deltaValues[dsId], float64(msg.Metadata.DeltaTime))

		if statsPerDataStructure[dsId].Oldest.Metadata.DeltaTime < msg.Metadata.DeltaTime {
			statsPerDataStructure[dsId].Oldest = msg
		}

		if statsPerDataStructure[dsId].Youngest.Metadata.DeltaTime > msg.Metadata.DeltaTime {
			statsPerDataStructure[dsId].Youngest = msg
		}

	}

	// fmt.Println("-----")

	for i, msStat := range statsPerDataStructure {
		msStat.MeanDeltaTime = stat.Mean(deltaValues[i], nil)
		// fmt.Printf("Id: %s, Mean: %f, Oldest: %+v, Young: %+v\n", i, statsPerDataStructure[i].MeanDeltaTime, statsPerDataStructure[i].Oldest, statsPerDataStructure[i].Youngest)
	}

	return statsPerDataStructure
}
