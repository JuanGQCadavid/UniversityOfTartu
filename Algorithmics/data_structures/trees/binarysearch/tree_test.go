package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dataValues = []int{
		1, 4, 2, 5, 3, 6, 9, 7, 8,
	}
	sorted = []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
)

func populateBST() *BSTrea[int, int] {
	bst := NewBSTree[int, int]()

	for _, data := range dataValues {
		bst.Append(data, data)
	}
	return bst
}

func TestPrint(t *testing.T) {
	bst := populateBST()
	bst.InOrderTreeWalk(bst.Root)
}

func TestSearch(t *testing.T) {
	bst := populateBST()

	node := bst.TreeSearch(bst.Root, dataValues[8])
	assert.NotEmpty(t, node)
	assert.Equal(t, node.Key, dataValues[8])
	assert.NotEqual(t, node.Key, -1)

	node2 := bst.IterativeTreeSearch(bst.Root, dataValues[8])
	assert.NotEmpty(t, node2)
	assert.Equal(t, node2.Key, dataValues[8])
	assert.NotEqual(t, node2.Key, -1)

	assert.Equal(t, node, node2)
}

func TestMax(t *testing.T) {
	bst := populateBST()

	node := bst.Maximun(bst.Root)
	assert.NotEmpty(t, node)
	assert.Equal(t, node.Key, sorted[len(sorted)-1])
	assert.NotEqual(t, node.Key, sorted[len(sorted)-2])
}

func TestMin(t *testing.T) {
	bst := populateBST()

	node := bst.Minimun(bst.Root)
	assert.NotEmpty(t, node)
	assert.Equal(t, node.Key, sorted[0])
	assert.NotEqual(t, node.Key, sorted[1])
}

func TestSuccessor(t *testing.T) {
	bst := populateBST()

	seven := bst.IterativeTreeSearch(bst.Root, 7)
	assert.NotEmpty(t, seven)

	six := bst.IterativeTreeSearch(bst.Root, 6)
	assert.NotEmpty(t, six)

	node := bst.Successor(six)
	print(node.SatalliteData)

	assert.NotEmpty(t, node)
	assert.Equal(t, seven, node)

	three := bst.IterativeTreeSearch(bst.Root, 3)
	assert.NotEmpty(t, seven)

	four := bst.IterativeTreeSearch(bst.Root, 4)
	assert.NotEmpty(t, six)

	node2 := bst.Successor(three)
	print(node2.SatalliteData)

	assert.NotEmpty(t, node2)
	assert.Equal(t, four, node2)

}
