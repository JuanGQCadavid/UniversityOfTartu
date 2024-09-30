package stats

import (
	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/internal/domain"
	"golang.org/x/exp/constraints"
)

func GetDepth[K constraints.Ordered, D any](node *domain.Node[K, D]) int {
	if node == nil {
		return 0
	}
	left, rigth := GetDepth(node.Left), GetDepth(node.Rigth)
	if left > rigth {
		return left + 1
	}
	return rigth + 1
}

func GetWidth[K constraints.Ordered, D any](node *domain.Node[K, D], onLevel int, levels map[int]int) int {
	if node == nil {
		return 0
	}

	levels[onLevel] = levels[onLevel] + GetWidth(node.Left, onLevel+1, levels) + GetWidth(node.Rigth, onLevel+1, levels)
	return 1
}
