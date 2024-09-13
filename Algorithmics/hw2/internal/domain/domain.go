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
	QueueId     string
}

type Msg struct {
	Id       string
	NextMsg  *Msg
	Metadata MessageMetadata
}

var (
	ErrNotItemsToPop = errors.New("not items left to pop")
)

type GeneralStats struct {
	Id           string
	DataType     OperationDataType
	ErrorsCount  int
	InsertCount  int
	DeleteCount  int
	MaxSizeCount int
	ActualSize   int
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
