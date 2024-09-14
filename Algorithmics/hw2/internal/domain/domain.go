package domain

import (
	"errors"
	"fmt"
)

type OperationDataType string

const (
	QUEUE OperationDataType = "Q"
	STACK OperationDataType = "S"
)

type OperationType string

const (
	INSERT OperationType = "INSERT"
	REMOVE OperationType = "REMOVE"
)

type Operation struct {
	DataType OperationDataType
	OpType   OperationType
	DatatId  string
	Timing   string
	Id       string
}

type MessageMetadata struct {
	TimeCreated string
	TimeDeleted string
	DeltaTime   int
	DSId        string
	DSType      OperationDataType
}

type Msg struct {
	Id       string
	NextMsg  *Msg
	Metadata MessageMetadata
}
type MsgStat struct {
	DataType      OperationDataType
	DataId        string
	MeanDeltaTime float64
	Oldest        *Msg
	Youngest      *Msg
}

var (
	ErrNotItemsToPop = errors.New("not items left to pop")
)

type GeneralStats struct {
	Id           string            `csv:"id"`
	DataType     OperationDataType `csv:"data_id"`
	ErrorsCount  int               `csv:"errors_count"`
	InsertCount  int               `csv:"insert_count"`
	DeleteCount  int               `csv:"delete_count"`
	MaxSizeCount int               `csv:"max_size"`
	ActualSize   int               `csv:"actual_size"`
}

func (gstats *GeneralStats) ToString() string {
	return fmt.Sprintf(
		"Stats: \n\tId: %s \n\tData type: %s \n\tErrorsCount: %06d \n\tInsertCount: %06d \n\tDeleteCount: %06d \n\tMaxSizeCount: %06d \n\tActualSize: %06d",
		gstats.Id,
		string(gstats.DataType),
		gstats.ErrorsCount,
		gstats.InsertCount,
		gstats.DeleteCount,
		gstats.MaxSizeCount,
		gstats.ActualSize,
	)
}

type GeneralCSV struct {
	TotalOperations int `csv:"total_operations"`
	TotalErrs       int `csv:"total_errs"`
	TotalInsert     int `csv:"total_insert"`
	TotalRemove     int `csv:"total_remove"`
}
