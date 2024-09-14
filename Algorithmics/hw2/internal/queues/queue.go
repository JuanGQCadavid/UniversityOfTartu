package queues

import (
	"hw2/internal/domain"
)

type Queue struct {
	ID        string
	actualMsg *domain.Msg
	records   []*domain.Msg
	errors    []error
	stats     domain.GeneralStats
}

func NewQueue(id string) *Queue {
	return &Queue{
		ID:        id,
		actualMsg: nil,
		records:   make([]*domain.Msg, 0),
		errors:    make([]error, 0),
		stats: domain.GeneralStats{
			Id:       id,
			DataType: domain.QUEUE,
		},
	}
}

func (queue *Queue) Pop(timeDeleted string) (*domain.Msg, error) {
	if queue.actualMsg == nil {
		queue.errors = append(queue.errors, domain.ErrNotItemsToPop)
		return nil, domain.ErrNotItemsToPop
	}
	queue.actualMsg.Metadata.TimeDeleted = timeDeleted     // Update Metadata
	oldMsg := queue.actualMsg                              // Store actual pointer to return
	queue.records = append(queue.records, queue.actualMsg) // Save it in our records
	queue.actualMsg = oldMsg.NextMsg                       // point to the older one
	queue.stats.DeleteCount++                              // Incrementing Stats
	return oldMsg, nil
}

func (queue *Queue) Push(msgId string, timeCreated string) {
	queue.stats.InsertCount++ // Incrementing Stats
	var nextMsg *domain.Msg = &domain.Msg{
		Id: msgId,
		Metadata: domain.MessageMetadata{
			TimeCreated: timeCreated,
			DSId:        queue.ID,
			DSType:      domain.QUEUE,
		},
	}

	if queue.stats.MaxSizeCount < (queue.stats.InsertCount - queue.stats.DeleteCount) {
		queue.stats.MaxSizeCount = (queue.stats.InsertCount - queue.stats.DeleteCount)
	}

	if queue.actualMsg == nil {
		// fmt.Println("I was empty, starting from ", nextMsg.Id)
		queue.actualMsg = nextMsg
	} else {
		lastElement := queue.actualMsg
		for {
			if lastElement.NextMsg != nil {
				lastElement = lastElement.NextMsg
			} else {
				break
			}
		}
		lastElement.NextMsg = nextMsg
		// fmt.Println("Last id ", lastElement.Id, " will point to", nextMsg.Id)

	}
}

func (queue *Queue) GetStats() domain.GeneralStats {
	return domain.GeneralStats{
		DataType:     queue.stats.DataType,
		Id:           queue.stats.Id,
		ErrorsCount:  len(queue.errors),
		ActualSize:   queue.stats.InsertCount - queue.stats.DeleteCount,
		InsertCount:  queue.stats.InsertCount,
		DeleteCount:  queue.stats.DeleteCount,
		MaxSizeCount: queue.stats.MaxSizeCount,
	}
}
