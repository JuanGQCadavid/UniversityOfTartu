package visualizer

import (
	"fmt"
	"testing"

	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/trees/binarysearch"
)

var (
	dataValues = []int{
		1, 4, 2, 5, 3, 6, 9, 7, 8,
	}
)

func populateBST() *binarysearch.BSTrea[int, int] {
	bst := binarysearch.NewBSTree[int, int]()

	for _, data := range dataValues {
		bst.Append(data, data)
	}
	return bst
}

func TestPrint(t *testing.T) {
	bst := populateBST()
	// log.Println(bst.NodeCounter(bst.Root))
	result := FromTreeToMatrix(bst.Root, bst.NodeCounter(bst.Root))

	for i := 0; i < len(result); i++ {
		fmt.Printf("%d - %+v\n", i, result[i])
	}
}

func TestFile(t *testing.T) {
	bst := populateBST()
	// log.Println(bst.NodeCounter(bst.Root))
	result := FromTreeToMatrix(bst.Root, bst.NodeCounter(bst.Root))

	for i := 0; i < len(result); i++ {
		fmt.Printf("%d - %+v\n", i, result[i])
	}

	FromMatrixToFile(result, "test.csv")
}

func TestPNG(t *testing.T) {
	bst := populateBST()
	result := FromTreeToMatrix(bst.Root, bst.NodeCounter(bst.Root))
	FromMatrixToFile(result, "test.csv")
	FromCSVToImage("test.csv")
}
