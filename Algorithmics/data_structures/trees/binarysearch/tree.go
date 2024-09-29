// Binary search tree property:  Let X be a node in a binary search tree.
// If Y is a node in the left subtree of x, then y.key <= key.
// If y is a node in the right subtree of x, then y.key >= key.
package binarysearch

import (
	"log"

	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/internal/domain"
	"golang.org/x/exp/constraints"
)

func NewBSTree[K constraints.Ordered, D any]() *BSTrea[K, D] {
	return &BSTrea[K, D]{
		Root: nil,
	}
}

type BSTrea[K constraints.Ordered, D any] struct {
	Root *domain.Node[K, D]
}

func (t *BSTrea[K, D]) Append(key K, data D) {
	if t.Root == nil {
		t.Root = &domain.Node[K, D]{
			Key:           key,
			SatalliteData: data,
		}
		return
	}

	actual := t.Root

	for {
		if key >= actual.Key {
			if actual.Rigth == nil {
				actual.Rigth = &domain.Node[K, D]{
					Key:           key,
					Parent:        actual,
					SatalliteData: data,
				}
				break
			}
			actual = actual.Rigth
		} else {
			if actual.Left == nil {
				actual.Left = &domain.Node[K, D]{
					Key:           key,
					Parent:        actual,
					SatalliteData: data,
				}
				break
			}
			actual = actual.Left
		}
	}

}

func (t *BSTrea[K, D]) InOrderTreeWalk(node *domain.Node[K, D]) {
	if node != nil {
		t.InOrderTreeWalk(node.Left)
		log.Printf("Data: %+v\n", node.SatalliteData)
		t.InOrderTreeWalk(node.Rigth)
	}
}

func (t *BSTrea[K, D]) NodeCounter(node *domain.Node[K, D]) int {
	if node != nil {
		return t.NodeCounter(node.Left) + 1 + t.NodeCounter(node.Rigth)
	}
	return 0
}

func (t *BSTrea[K, D]) PreOrderTreeWalk(node *domain.Node[K, D]) {
	if node != nil {
		log.Printf("Data: %+v\n", node.SatalliteData)
		t.InOrderTreeWalk(node.Left)
		t.InOrderTreeWalk(node.Rigth)
	}
}

func (t *BSTrea[K, D]) PostOrderTreeWalk(node *domain.Node[K, D]) {
	if node != nil {
		t.InOrderTreeWalk(node.Left)
		t.InOrderTreeWalk(node.Rigth)
		log.Printf("Data: %+v\n", node.SatalliteData)
	}
}

func (t *BSTrea[K, D]) TreeSearch(node *domain.Node[K, D], key K) *domain.Node[K, D] {
	if node == nil || node.Key == key {
		return node
	}

	if key < node.Key {
		return t.TreeSearch(node.Left, key)
	}
	return t.TreeSearch(node.Rigth, key)
}

func (t *BSTrea[K, D]) IterativeTreeSearch(node *domain.Node[K, D], key K) *domain.Node[K, D] {
	for node != nil && node.Key != key {
		if key < node.Key {
			node = node.Left
			continue
		}
		node = node.Rigth
	}
	return node
}

func (t *BSTrea[K, D]) Minimun(node *domain.Node[K, D]) *domain.Node[K, D] {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (t *BSTrea[K, D]) Maximun(node *domain.Node[K, D]) *domain.Node[K, D] {
	for node.Rigth != nil {
		node = node.Rigth
	}
	return node
}

func (t *BSTrea[K, D]) Successor(node *domain.Node[K, D]) *domain.Node[K, D] {

	if node.Rigth != nil {
		return t.Minimun(node.Rigth)
	}

	y := node.Parent

	for y != nil && node == y.Rigth {
		node = y
		y = y.Parent
	}

	return y
}
