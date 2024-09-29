package stacks

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestFailPop(t *testing.T) {
	stack := NewStack("1")

	if _, err := stack.Pop("X"); err != domain.ErrNotItemsToPop {
		t.Error("Different error was returned, err: ", err.Error(), "expecting: ", domain.ErrNotItemsToPop.Error())
	}

	stack.Push("X", "X")

	if pop, err := stack.Pop("X"); err != nil {
		t.Error("Error was not expected but it was returned, err: ", err.Error())
	} else if pop.Id != "X" {
		t.Error("Id was not Expected, ", pop.Id)
	}

	if _, err := stack.Pop("X"); err != domain.ErrNotItemsToPop {
		t.Error("Different error was returned, err: ", err.Error(), "expecting: ", domain.ErrNotItemsToPop.Error())
	}
}

// type GeneralStats struct {
// 	ErrorsCount  int
// 	InsertCount  int
// 	DeleteCount  int
// 	MaxSizeCount int
// 	ActualSize   int
// }

func TestStats(t *testing.T) {
	var tests = []struct {
		name            string
		operations      []domain.Operation
		insertCount     int
		deletecount     int
		maxSizeCount    int
		actualSizeCount int
		errorsCounts    int
	}{
		{
			name: "Insert count check",
			operations: []domain.Operation{
				{
					OpType: domain.INSERT,
				},
				{
					OpType: domain.INSERT,
				},
				{
					OpType: domain.INSERT,
				},
				{
					OpType: domain.INSERT,
				},
				{
					OpType: domain.INSERT,
				},
			},
			insertCount:     5,
			deletecount:     0,
			maxSizeCount:    5,
			actualSizeCount: 5,
			errorsCounts:    0,
		},
		{
			name: "Insert and Delete count check",
			operations: []domain.Operation{
				{
					OpType: domain.INSERT,
				},
				{
					OpType: domain.REMOVE,
				},
				{
					OpType: domain.INSERT,
				},
				{
					OpType: domain.REMOVE,
				},
				{
					OpType: domain.INSERT,
				},
			},
			insertCount:     3,
			deletecount:     2,
			maxSizeCount:    1,
			actualSizeCount: 1,
			errorsCounts:    0,
		},
		{
			name: "Max Count count check",
			operations: []domain.Operation{
				{
					OpType: domain.INSERT,
				},
				{
					OpType: domain.REMOVE,
				},
				{
					OpType: domain.INSERT,
				},
				{
					OpType: domain.INSERT,
				},
				{
					OpType: domain.INSERT,
				},
				{
					OpType: domain.REMOVE,
				},
				{
					OpType: domain.INSERT,
				},
			},
			insertCount:     5,
			deletecount:     2,
			maxSizeCount:    3,
			actualSizeCount: 3,
			errorsCounts:    0,
		},
		{
			name: "Actual Size check",
			operations: []domain.Operation{
				{
					OpType: domain.INSERT,
				},
				{
					OpType: domain.REMOVE,
				},
				{
					OpType: domain.INSERT,
				},
				{
					OpType: domain.REMOVE,
				},
			},
			insertCount:     2,
			deletecount:     2,
			maxSizeCount:    1,
			actualSizeCount: 0,
			errorsCounts:    0,
		},
		{
			name: "Erros check",
			operations: []domain.Operation{
				{
					OpType: domain.INSERT,
				},
				{
					OpType: domain.REMOVE,
				},
				{
					OpType: domain.REMOVE,
				},
				{
					OpType: domain.INSERT,
				},
				{
					OpType: domain.REMOVE,
				},
				{
					OpType: domain.REMOVE,
				},
			},
			insertCount:     2,
			deletecount:     2,
			maxSizeCount:    1,
			actualSizeCount: 0,
			errorsCounts:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack("1")
			for index, op := range tt.operations {
				if op.OpType == domain.INSERT {
					stack.Push(strconv.Itoa(index), fmt.Sprintf("T%06d", index))
				} else {
					stack.Pop(fmt.Sprintf("T%06d", index))
				}
			}
			statsGenerated := stack.GetStats()
			assert.Equal(t, tt.actualSizeCount, statsGenerated.ActualSize, "Actual size check")
			assert.Equal(t, tt.insertCount, statsGenerated.InsertCount, "Insert count check")
			assert.Equal(t, tt.deletecount, statsGenerated.DeleteCount, "Delete count check")
			assert.Equal(t, tt.maxSizeCount, statsGenerated.MaxSizeCount, "Max count check")
			assert.Equal(t, tt.actualSizeCount, statsGenerated.ActualSize, "Actual Size check")
			assert.Equal(t, tt.errorsCounts, statsGenerated.ErrorsCount, "Insert count check")
		})
	}
}

func TestPopPush(t *testing.T) {
	var tests = []struct {
		name string
		n    int
	}{
		{
			name: "One Push Pop",
			n:    1,
		},
		{
			name: "Ten Push Pop",
			n:    10,
		},
		{
			name: "One Thousand Push Pop",
			n:    1000,
		},
		{
			name: "Then Thousand Push Pop",
			n:    10000,
		},
		{
			name: "One houndred Thousand Push Pop",
			n:    100000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack("1")

			messages := make([]domain.Operation, tt.n)
			returnedMessages := make([]bool, tt.n)

			for index := range messages {
				messages[index] = domain.Operation{
					Id:     strconv.Itoa(index),
					Timing: fmt.Sprintf("T%06d", index),
				}
			}

			for _, msg := range messages {
				// t.Logf("Saving Id %s, time %s\n", msg.Id, msg.Timing)
				stack.Push(msg.Id, msg.Timing)
			}

			for index := range messages {
				msgReturned, err := stack.Pop(fmt.Sprintf("T%06d", index+len(messages)))
				// t.Logf("msgReturned: Id %s, time Added %s, time Removed %s, Queue Id: %s\n",
				// 	msgReturned.Id,
				// 	msgReturned.Metadata.TimeCreated,
				// 	msgReturned.Metadata.TimeDeleted,
				// 	msgReturned.Metadata.QueueId,
				// )

				if err != nil {
					t.Error("An error was not expected. Err: ", err.Error())
				}

				if messages[len(messages)-(index+1)].Id != msgReturned.Id {
					t.Error("Pop returns ", msgReturned.Id, " I was waiting for", messages[len(messages)-(index+1)].Id)
				}

				for ii, msgs := range messages {
					if msgs.Id == msgReturned.Id {
						returnedMessages[ii] = true
					}
				}
			}

			for index := range returnedMessages {
				if !returnedMessages[index] {
					t.Error("Message Id", index, "Was not returned")
				}
			}

		})
	}
}
