package ports

import "hw2/internal/domain"

type DataStructure interface {
	Push(msgId string, timeCreated string)
	Pop(timeDeleted string) (*domain.Msg, error)
	GetStats() domain.GeneralStats
}
