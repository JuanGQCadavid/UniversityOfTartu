package btree

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBTreeInsertRoot(t *testing.T) {
	var (
		maxValuesPerNode = 4
		values           = []int{
			5, 8, 9, 10,
		}
		btree = NewBTree[int](maxValuesPerNode)
	)

	for _, val := range values {
		btree.InsertValue(val, btree.Root)
	}

	btree.PrintTree()
	assert.Equal(t, btree.Root.Values, values)
}

func TestBTreeRootCreated(t *testing.T) {
	var (
		maxValuesPerNode = 4
		values           = []int{
			5, 8, 9, 10, 11,
		}
		btree = NewBTree[int](maxValuesPerNode)
	)

	for _, val := range values {
		btree.InsertValue(val, btree.Root)
	}

	btree.PrintTree()
	assert.Equal(t, btree.Root.Values, []int{9})
	assert.Equal(t, btree.Root.Pointers[0].Values, []int{5, 8})
	assert.Equal(t, btree.Root.Pointers[1].Values, []int{10, 11})

}

func TestBTreeInsert(t *testing.T) {
	var (
		maxValuesPerNode = 4
		values           = []int{
			5, 8, 9, 10, 11, 12, 4, 6, 13, 16,
		}
		btree = NewBTree[int](maxValuesPerNode)
	)
	for _, val := range values {
		log.Println("To insert", val)
		btree.InsertValue(val, btree.Root)
		btree.PrintTree()
	}

	assert.Equal(t, btree.Root.Values, []int{9, 12})
	assert.Equal(t, btree.Root.Pointers[0].Values, []int{4, 5, 6, 8})
	assert.Equal(t, btree.Root.Pointers[1].Values, []int{10, 11})
	assert.Equal(t, btree.Root.Pointers[2].Values, []int{13, 16})
}

func TestBTreeInsert2Levels4Buckets(t *testing.T) {
	var (
		maxValuesPerNode = 4
		values           = []int{
			5, 8, 9, 10, 11, 12, 4, 6, 13, 16, 7,
		}
		btree = NewBTree[int](maxValuesPerNode)
	)
	for _, val := range values {
		log.Println("To insert", val)
		btree.InsertValue(val, btree.Root)
		btree.PrintTree()
	}

	assert.Equal(t, btree.Root.Values, []int{6, 9, 12})
	assert.Equal(t, btree.Root.Pointers[0].Values, []int{4, 5})
	assert.Equal(t, btree.Root.Pointers[1].Values, []int{7, 8})
	assert.Equal(t, btree.Root.Pointers[2].Values, []int{10, 11})
	assert.Equal(t, btree.Root.Pointers[3].Values, []int{13, 16})
}

func TestBTreeInsert3Levels(t *testing.T) {
	var (
		maxValuesPerNode = 4
		values           = []int{
			5, 8, 9, 10, 11, 12, 4, 6, 13, 16, 7, 1, 3, 2, 14, 18, 20,
		}
		btree = NewBTree[int](maxValuesPerNode)
	)
	for _, val := range values {
		log.Println("To insert", val)
		btree.InsertValue(val, btree.Root)
		btree.PrintTree()
	}

	// assert.Equal(t, btree.Root.Values, []int{6, 9, 12})
	// assert.Equal(t, btree.Root.Pointers[0].Values, []int{4, 5})
	// assert.Equal(t, btree.Root.Pointers[1].Values, []int{7, 8})
	// assert.Equal(t, btree.Root.Pointers[2].Values, []int{10, 11})
	// assert.Equal(t, btree.Root.Pointers[3].Values, []int{13, 16})
}
