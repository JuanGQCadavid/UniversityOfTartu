package search

import (
	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/internal/domain"
	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/queues"
	"golang.org/x/exp/constraints"
)

func BreadthFirstSearchAll[K constraints.Ordered, D any](root *domain.Node[K, D], count int) []D {
	results := make([]D, 0, count)

	pendingToVisit := queues.NewQueue[*domain.Node[K, D]]("PendingToVisit")
	pendingToVisit.Push(root)

	for {
		nodeToVisit, err := pendingToVisit.Pop()

		if err == domain.ErrNotItemsToPop {
			break
		}

		results = append(results, nodeToVisit.Value.SatalliteData)

		if nodeToVisit.Value.Left != nil {
			pendingToVisit.Push(nodeToVisit.Value.Left)
		}
		if nodeToVisit.Value.Rigth != nil {
			pendingToVisit.Push(nodeToVisit.Value.Rigth)
		}
	}

	return results

}

func MaximunPathSum[K constraints.Ordered, D constraints.Integer | constraints.Float](root *domain.Node[K, D]) (D, []*domain.Node[K, D]) {
	if root == nil {
		return 0, []*domain.Node[K, D]{}
	}

	leftCount, leftNodes := MaximunPathSum(root.Left)
	rigthCount, rigthNodes := MaximunPathSum(root.Rigth)

	if leftCount > rigthCount {
		leftNodes = append(leftNodes, root)
		leftCount = leftCount + root.SatalliteData
		return leftCount, leftNodes
	}

	rigthNodes = append(rigthNodes, root)
	rigthCount = rigthCount + root.SatalliteData
	return rigthCount, rigthNodes
}
