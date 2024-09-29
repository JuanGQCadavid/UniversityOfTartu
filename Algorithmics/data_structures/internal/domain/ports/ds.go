package ports

import "github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/internal/domain"

type DataStructure interface {
	Push(msgId string, timeCreated string)
	Pop(timeDeleted string) (*domain.Msg, error)
	GetStats() domain.GeneralStats
}
