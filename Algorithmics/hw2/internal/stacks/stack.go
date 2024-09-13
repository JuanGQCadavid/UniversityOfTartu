package stacks

import (
	"hw2/internal/domain"
)

type Stack struct {
	ID        string
	actualMsg *domain.Msg
	records   []*domain.Msg
	errors    []error
}

func NewStack(id string) *Stack {
	return &Stack{
		ID:        id,
		actualMsg: nil,
		records:   make([]*domain.Msg, 0),
		errors:    make([]error, 0),
	}
}

func (stack *Stack) Pop(timeDeleted string) (*domain.Msg, error) {
	if stack.actualMsg == nil {
		stack.errors = append(stack.errors, domain.ErrNotItemsToPop)
		return nil, domain.ErrNotItemsToPop
	}

	stack.actualMsg.Metadata.TimeDeleted = timeDeleted // Update Metadata

	oldMsg := stack.actualMsg

	stack.records = append(stack.records, stack.actualMsg) // Save it in our records
	stack.actualMsg = stack.actualMsg.NextMsg              // point to the older one
	return oldMsg, nil
}

func (stack *Stack) Push(msgId string, timeCreated string) {
	stack.actualMsg = &domain.Msg{
		NextMsg: stack.actualMsg,
		Id:      msgId,
		Metadata: domain.MessageMetadata{
			TimeCreated: timeCreated,
			QueueId:     stack.ID,
		},
	}
}
