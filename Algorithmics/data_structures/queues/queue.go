package queues

import (
	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/internal/domain"
)

type Msg[D any] struct {
	Value   D
	NextMsg *Msg[D]
}

type Queue[D any] struct {
	ID        string
	actualMsg *Msg[D]
}

func NewQueue[D any](id string) *Queue[D] {
	return &Queue[D]{
		ID:        id,
		actualMsg: nil,
	}
}

func (queue *Queue[D]) Pop() (*Msg[D], error) {
	if queue.actualMsg == nil {
		return nil, domain.ErrNotItemsToPop
	}
	oldMsg := queue.actualMsg        // Store actual pointer to return
	queue.actualMsg = oldMsg.NextMsg // point to the older one
	return oldMsg, nil
}

func (queue *Queue[D]) Push(value D) {
	var nextMsg *Msg[D] = &Msg[D]{
		Value: value,
	}

	if queue.actualMsg == nil {
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
	}
}
