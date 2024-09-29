package btree

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"golang.org/x/exp/constraints"
)

type BNode[K constraints.Ordered] struct {
	Values   []K
	Pointers []*BNode[K]
	Parent   *BNode[K]
}

type BTree[K constraints.Ordered] struct {
	Root            *BNode[K]
	ValuesPerNode   int
	PointersPerNode int
	MaxValues       int
	MinValues       int
}

func NewBTree[K constraints.Ordered](valuesPerNode int) *BTree[K] {
	return &BTree[K]{
		Root:            nil,
		ValuesPerNode:   valuesPerNode,
		PointersPerNode: valuesPerNode + 1,
		MaxValues:       valuesPerNode,
		MinValues:       valuesPerNode / 2,
	}
}

// PrintTree recursively prints the B-tree structure
func (tree *BTree[K]) PrintTree() {
	if tree.Root == nil {
		fmt.Println("Tree is empty")
		return
	}
	tree.printNode(tree.Root, 0)
}

// Helper function to print each node
func (tree *BTree[K]) printNode(node *BNode[K], depth int) {
	if node == nil {
		return
	}

	// Print indentation for tree depth
	indent := strings.Repeat("  ", depth)

	// Print node values
	fmt.Printf("%sNode Values: %v\n", indent, node.Values)

	// Recursively print child nodes (pointers)
	for i, pointer := range node.Pointers {
		if pointer != nil {
			fmt.Printf("%s Pointer %d:\n", indent, i)
			tree.printNode(pointer, depth+1)
		}
	}
}

func (tree *BTree[K]) printNode2(node *BNode[K], depth int) {
	if node == nil {
		return
	}

	// Print indentation for tree depth
	indent := strings.Repeat("  ", depth)

	// Print node values
	fmt.Printf("%sNode Values: %v\n", indent, node.Values)

	// Prepare to print child nodes (pointers) on the same line
	if len(node.Pointers) > 0 {
		fmt.Printf("%sChild Pointers:\n", indent)
		fmt.Printf("%s", indent)

		// Print child node values inline
		for _, pointer := range node.Pointers {
			if pointer != nil {
				fmt.Printf(" [%v] ", pointer.Values)
			} else {
				fmt.Printf(" [nil] ")
			}
		}
		fmt.Println()

		// Recursively print each child node
		for _, pointer := range node.Pointers {
			tree.printNode2(pointer, depth+1)
		}
	}
}

// Rules:

// 1. All Leaves should be at the same level
// (The leaves are the nodes that are at the bottom, whose Pointers are to null)

// 2. Every node has a maximun and a minimun number of keys
// Max: Input
// Min: Half of the Max number of values
// Exeption: The root node, it is the only one who can have less than the min

// Proces:

// Every time we add a key it is added in the arr in sorted mode
// We Split a node when the array that will contained would have more than the max
func (btree *BTree[K]) InsertValue(value K, node *BNode[K]) {
	if btree.Root == nil {
		// Case 1. Root is nill

		btree.Root = &BNode[K]{
			Values:   make([]K, 0, btree.ValuesPerNode),
			Pointers: make([]*BNode[K], 0, btree.PointersPerNode),
		}
		btree.Root.Values = append(btree.Root.Values, value)
		return
	}
	// Insert
	// I have pointers, then try to fit there

	sum := 0
	for _, pointer := range node.Pointers {
		if pointer != nil {
			sum += 1
		}
	}

	if sum > 0 {
		insertIndex := sort.Search(len(node.Values), func(i int) bool { return value < node.Values[i] })
		// index := len(node.Values)
		// for i, val := range node.Values {
		// 	if value < val {
		// 		index = i
		// 	}
		// }
		// Lets go to the node That would have me
		btree.InsertValue(value, node.Pointers[insertIndex])
		return
	} else if len(node.Values) < btree.MaxValues { // Well, It seems I'm the bottom, could I hold the new value?
		// Yes! I have space, just add it here in sorted way
		node.Values = append(node.Values, value)
		sort.Slice(node.Values, func(i, j int) bool {
			return node.Values[i] < node.Values[j]
		})
		return
	} else {
		// Well, the value should be here, but I dont have space
		// Time for revalancing
		log.Println("Shit, i dont know how to rebalance, lets try")
		btree.Rebalancing(value, node)
	}
}

// func (btree *BTree[K]) findPointerPosition(value K, node *BNode[K]) int {

// }

func (btree *BTree[K]) NewBNode() *BNode[K] {
	newPointer := make([]*BNode[K], 0, btree.PointersPerNode)

	for i := 0; i < btree.PointersPerNode; i++ {
		newPointer = append(newPointer, nil)
	}

	return &BNode[K]{
		Values:   make([]K, 0, btree.MaxValues),
		Pointers: newPointer,
	}
}

func (btree *BTree[K]) Rebalancing(value K, node *BNode[K]) {
	log.Printf("Rebalancing: %+v", node)
	if node.Parent != nil {
		log.Println(" Parent: ", node.Parent.Values)
	}
	newNode := btree.NewBNode()

	valuesToSplit := make([]K, 0, btree.MaxValues+1)
	valuesToSplit = append(valuesToSplit, node.Values...)
	valuesToSplit = append(valuesToSplit, value)

	// Sorting in order to matain b tree order
	sort.Slice(valuesToSplit, func(i, j int) bool {
		return valuesToSplit[i] < valuesToSplit[j]
	})

	midIndex := len(valuesToSplit) / 2
	node.Values = valuesToSplit[0:midIndex]
	midValue := valuesToSplit[midIndex]
	newNode.Values = valuesToSplit[midIndex+1:]

	log.Println("OldNode: ", node.Values, " Mid: ", midValue, " NewNode: ", newNode.Values)

	// Manage pointers: move half of the pointers to the new node if it's an internal node
	if len(node.Pointers) > 0 {
		placeHolder := make([]*BNode[K], 0, btree.PointersPerNode)

		for i := 0; i < btree.PointersPerNode; i++ {
			placeHolder = append(placeHolder, nil)
		}

		newNode.Pointers = append(node.Pointers[midIndex+1:], placeHolder...)[:btree.PointersPerNode]
		node.Pointers = append(node.Pointers[:midIndex+1], placeHolder...)[:btree.PointersPerNode] //node.Pointers[:midIndex+1]
	}

	// Well, it seems we are on the root, so a new root would be created
	if node.Parent == nil {
		log.Println("New root")
		newRoot := btree.NewBNode()

		// Update parent pointer
		node.Parent = newRoot
		newNode.Parent = newRoot

		newRoot.Values = append(newRoot.Values, midValue)
		newRoot.Pointers[0] = node    // Left
		newRoot.Pointers[1] = newNode // Rigth

		log.Printf("Node: %+v", node)
		log.Printf("New Node: %+v", newNode)

		// Update root pointer
		btree.Root = newRoot
		return
	}

	// Well, the parent should either have it or split itself
	parent := node.Parent

	// Does the root have space to insert a new value?
	if len(parent.Values) < btree.MaxValues { // Yes it has!
		log.Println("*****************************************************Hola?")
		// Mid index
		insertIndex := sort.Search(len(parent.Values), func(i int) bool { return midValue < parent.Values[i] })

		// Parent new values
		parent.Values = append(parent.Values[:insertIndex], append([]K{midValue}, parent.Values[insertIndex:]...)...)

		// Parent new pointes
		parent.Pointers = append(parent.Pointers[:insertIndex+1], append([]*BNode[K]{newNode}, parent.Pointers[insertIndex+1:]...)...)[:btree.PointersPerNode]

		// New node parent addition
		newNode.Parent = node.Parent
	} else {
		parent.Pointers = append(parent.Pointers, newNode)
		newNode.Parent = node.Parent
		log.Println("Parent is full, rebalance needed")
		btree.Rebalancing(midValue, parent)
	}
}
