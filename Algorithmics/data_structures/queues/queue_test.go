package queues

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestStats(t *testing.T) {
	var tests = []struct {
		name          string
		operations    []domain.Operation
		statsExpected domain.GeneralStats
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
			statsExpected: domain.GeneralStats{
				InsertCount:  5,
				DeleteCount:  0,
				MaxSizeCount: 5,
				ActualSize:   5,
				ErrorsCount:  0,
			},
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
			statsExpected: domain.GeneralStats{
				InsertCount:  3,
				DeleteCount:  2,
				MaxSizeCount: 1,
				ActualSize:   1,
				ErrorsCount:  0,
			},
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
			statsExpected: domain.GeneralStats{
				InsertCount:  5,
				DeleteCount:  2,
				MaxSizeCount: 3,
				ActualSize:   3,
				ErrorsCount:  0,
			},
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
			statsExpected: domain.GeneralStats{
				InsertCount:  2,
				DeleteCount:  2,
				MaxSizeCount: 1,
				ActualSize:   0,
				ErrorsCount:  0,
			},
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
			statsExpected: domain.GeneralStats{
				InsertCount:  2,
				DeleteCount:  2,
				MaxSizeCount: 1,
				ActualSize:   0,
				ErrorsCount:  2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewQueue("1")
			for index, op := range tt.operations {
				if op.OpType == domain.INSERT {
					queue.Push(strconv.Itoa(index), fmt.Sprintf("T%06d", index))
				} else {
					queue.Pop(fmt.Sprintf("T%06d", index))
				}
			}
			statsGenerated := queue.GetStats()
			assert.Equal(t, tt.statsExpected.InsertCount, statsGenerated.InsertCount, "Insert count check")
			assert.Equal(t, tt.statsExpected.DeleteCount, statsGenerated.DeleteCount, "Delete count check")
			assert.Equal(t, tt.statsExpected.MaxSizeCount, statsGenerated.MaxSizeCount, "Max count check")
			assert.Equal(t, tt.statsExpected.ActualSize, statsGenerated.ActualSize, "Actual Size check")
			assert.Equal(t, tt.statsExpected.ErrorsCount, statsGenerated.ErrorsCount, "Insert count check")
		})
	}
}

func TestFailPop(t *testing.T) {
	queue := NewQueue("1")

	if _, err := queue.Pop("X"); err != domain.ErrNotItemsToPop {
		t.Error("Different error was returned, err: ", err.Error(), "expecting: ", domain.ErrNotItemsToPop.Error())
	}

	queue.Push("X", "X")

	if pop, err := queue.Pop("X"); err != nil {
		t.Error("Error was not expected but it was returned, err: ", err.Error())
	} else if pop.Id != "X" {
		t.Error("Id was not Expected, ", pop.Id)
	}

	if _, err := queue.Pop("X"); err != domain.ErrNotItemsToPop {
		t.Error("Different error was returned, err: ", err.Error(), "expecting: ", domain.ErrNotItemsToPop.Error())
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewQueue("1")

			messages := make([]domain.Operation, tt.n)
			returnedMessages := make([]bool, tt.n)

			for index := range messages {
				messages[index] = domain.Operation{
					Id:     strconv.Itoa(index),
					Timing: fmt.Sprintf("T%06d", index),
				}
			}

			for _, msg := range messages {
				t.Logf("Saving Id %s, time %s\n", msg.Id, msg.Timing)
				queue.Push(msg.Id, msg.Timing)
			}

			for index, _ := range messages {
				msgReturned, err := queue.Pop(fmt.Sprintf("T%06d", index+len(messages)))

				if err != nil {
					t.Fatal("An error was not expected. Err: ", err.Error())
				}

				t.Logf("msgReturned: Id %s, time Added %s, time Removed %s, Queue Id: %s\n",
					msgReturned.Id,
					msgReturned.Metadata.TimeCreated,
					msgReturned.Metadata.TimeDeleted,
					msgReturned.Metadata.DSId,
				)

				if msgReturned == nil {
					t.Error("msgReturned is nill")
				}

				if messages[index].Id != msgReturned.Id {
					t.Error("Pop returns ", msgReturned.Id, " I was waiting for", messages[index].Id)
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
