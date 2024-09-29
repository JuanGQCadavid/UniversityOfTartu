package search

import (
	"fmt"
	"log"
	"testing"

	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/internal/domain"
	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/trees/binarysearch"
	"github.com/stretchr/testify/assert"
)

var (
	dataValues = []int{
		1, 4, 2, 5, 3, 6, 9, 7, 8,
	}
	bfs = []int{
		1, 4, 2, 5, 3, 6, 9, 7, 8,
	}
	longDataValues = []int{
		51, 154, 77, 232, 116, 58, 29, 88, 44, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2,
	}
	bfsLong = []int{
		51, 29, 154, 22, 44, 77, 232, 11, 26, 34, 58, 116, 10, 17, 40, 52, 88, 5, 13, 20, 4, 8, 16, 2,
	}
)

func populateBST(dataValues []int) *binarysearch.BSTrea[int, int] {
	bst := binarysearch.NewBSTree[int, int]()

	for _, data := range dataValues {
		bst.Append(data, data)
	}
	return bst
}

func TestBFS(t *testing.T) {
	bst := populateBST(dataValues)
	result := BreadthFirstSearchAll(bst.Root, bst.NodeCounter(bst.Root))
	for i := range bfs {
		assert.Equal(t, bfs[i], result[i])
	}
}

func TestBFS2(t *testing.T) {
	bst := populateBST(longDataValues)
	result := BreadthFirstSearchAll(bst.Root, bst.NodeCounter(bst.Root))
	for i := range bfsLong {
		assert.Equal(t, bfsLong[i], result[i])
	}
}

func TestMaximunPathSum(t *testing.T) {
	bst := populateBST(longDataValues)
	sumL, nodesL := MaximunPathSum(bst.Root.Left)
	sumR, nodesR := MaximunPathSum(bst.Root.Rigth)
	log.Println("sumL:", sumL)
	log.Println("sumR:", sumR)

	// for _, node := range nodesL {
	// 	log.Println(node.SatalliteData)
	// }
	// log.Println("(")
	// log.Println(bst.Root.SatalliteData)
	// log.Println(")")
	// for _, node := range nodesR {
	// 	log.Println(node.SatalliteData)
	// }

	totalNodes := make([]*domain.Node[int, int], 0, len(nodesL)+len(nodesR)+1)

	totalNodes = append(totalNodes, nodesL...)
	totalNodes = append(totalNodes, bst.Root)
	totalNodes = append(totalNodes, nodesR...)
	newPointer := len(totalNodes) - 1
	for i, node := range totalNodes {
		if i == len(nodesL) {
			fmt.Print("(", node.SatalliteData, "), ")
			continue
		}

		if i > len(nodesL) {
			// pos := (len(nodesL) - 1 + len(nodesR) - 1) - (len(nodesL) - 1) + i
			fmt.Print(totalNodes[newPointer].SatalliteData, ", ")
			newPointer--
			continue
		}

		fmt.Print(node.SatalliteData, ", ")
	}

}
