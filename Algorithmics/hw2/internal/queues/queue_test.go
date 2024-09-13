package queues

import (
	"fmt"
	"hw2/internal/domain"
	"strconv"
	"testing"
)

func TestFailPop(t *testing.T) {
	stack := NewQueue("1")

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
					msgReturned.Metadata.QueueId,
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
