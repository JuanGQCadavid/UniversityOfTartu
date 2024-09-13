package domain

import "errors"

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
