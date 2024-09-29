package btree

import (
	"log"
	"testing"
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

}
