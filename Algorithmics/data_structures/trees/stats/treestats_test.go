package stats

import (
	"log"
	"testing"

	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/trees/binarysearch"
	"github.com/stretchr/testify/assert"
)

var (
	dataValues = []int{
		1, 4, 2, 5, 3, 6, 9, 7, 8,
	}
	longDataValues = []int{
		51, 154, 77, 232, 116, 58, 29, 88, 44, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2,
	}
	width = map[int]int{
		7: 1,
		6: 3,
		5: 3,
		4: 5,
		3: 5,
		2: 4,
		1: 2,
		8: 0,
	}
)

func populateBST(dataValues []int) *binarysearch.BSTrea[int, int] {
	bst := binarysearch.NewBSTree[int, int]()

	for _, data := range dataValues {
		bst.Append(data, data)
	}
	return bst
}
func TestDepth(t *testing.T) {
	bst := populateBST(dataValues)
	depth := GetDepth(bst.Root)
	log.Println(depth)
	assert.Equal(t, 7, depth)
}

func TestWidth(t *testing.T) {
	bst := populateBST(longDataValues)
	result := make(map[int]int)
	GetWidth(bst.Root, 1, result)
	for key := range result {
		assert.Equal(t, width[key], result[key])
		log.Println(key, result[key])
	}
}
