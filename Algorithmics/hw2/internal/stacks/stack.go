package stacks

import (
	"hw2/internal/domain"
)

type Stack struct {
	ID        string
	actualMsg *domain.Msg
	records   []*domain.Msg
	errors    []error
	stats     domain.GeneralStats
}

func NewStack(id string) *Stack {
	return &Stack{
		ID:        id,
		actualMsg: nil,
		records:   make([]*domain.Msg, 0),
		errors:    make([]error, 0),
		stats: domain.GeneralStats{
			Id:       id,
			DataType: domain.STACK,
		},
	}
}

func (stack *Stack) Pop(timeDeleted string) (*domain.Msg, error) {
	if stack.actualMsg == nil {
		stack.errors = append(stack.errors, domain.ErrNotItemsToPop)
		return nil, domain.ErrNotItemsToPop
	}
	stack.actualMsg.Metadata.TimeDeleted = timeDeleted     // Update Metadata
	oldMsg := stack.actualMsg                              // Saving the actual MSG
	stack.records = append(stack.records, stack.actualMsg) // Save it in our records
	stack.actualMsg = stack.actualMsg.NextMsg              // point to the older one
	stack.stats.DeleteCount++                              // Incrementing Stats
	return oldMsg, nil
}

func (stack *Stack) Push(msgId string, timeCreated string) {
	stack.stats.InsertCount++ // Incrementing Stats
	stack.actualMsg = &domain.Msg{
		NextMsg: stack.actualMsg,
		Id:      msgId,
		Metadata: domain.MessageMetadata{
			TimeCreated: timeCreated,
			QueueId:     stack.ID,
		},
	}

	if stack.stats.MaxSizeCount < (stack.stats.InsertCount - stack.stats.DeleteCount) {
		stack.stats.MaxSizeCount = (stack.stats.InsertCount - stack.stats.DeleteCount)
	}
}

func (stack *Stack) GetStats() domain.GeneralStats {
	return domain.GeneralStats{
		Id:           stack.stats.Id,
		DataType:     stack.stats.DataType,
		ErrorsCount:  len(stack.errors),
		ActualSize:   stack.stats.InsertCount - stack.stats.DeleteCount,
		InsertCount:  stack.stats.InsertCount,
		DeleteCount:  stack.stats.DeleteCount,
		MaxSizeCount: stack.stats.MaxSizeCount,
	}
}
